// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package datastore

import (
	"fmt"
	"reflect"
	"time"

	pb "google.golang.org/cloud/internal/datastore"
)

var (
	typeOfBlobKey   = reflect.TypeOf(BlobKey(""))
	typeOfByteSlice = reflect.TypeOf([]byte(nil))
	typeOfTime      = reflect.TypeOf(time.Time{})
)

// typeMismatchReason returns a string explaining why the property p could not
// be stored in an entity field of type v.Type().
func typeMismatchReason(p Property, v reflect.Value) string {
	entityType := "empty"
	switch p.Value.(type) {
	case int64:
		entityType = "int"
	case bool:
		entityType = "bool"
	case string:
		entityType = "string"
	case float64:
		entityType = "float"
	case *Key:
		entityType = "*datastore.Key"
	case time.Time:
		entityType = "time.Time"
	case BlobKey:
		entityType = "datastore.BlobKey"
	case []byte:
		entityType = "[]byte"
	}
	return fmt.Sprintf("type mismatch: %s versus %v", entityType, v.Type())
}

type propertyLoader struct {
	// m holds the number of times a substruct field like "Foo.Bar.Baz" has
	// been seen so far. The map is constructed lazily.
	m map[string]int
}

func (l *propertyLoader) load(codec *structCodec, structValue reflect.Value, p Property, prev map[string]struct{}) string {
	var sliceOk bool
	var v reflect.Value
	// Traverse a struct's struct-typed fields.
	for name := p.Name; ; {
		decoder, ok := codec.byName[name]
		if !ok {
			return "no such struct field"
		}
		v = structValue.Field(decoder.index)
		if !v.IsValid() {
			return "no such struct field"
		}
		if !v.CanSet() {
			return "cannot set struct field"
		}

		if decoder.substructCodec == nil {
			break
		}

		if v.Kind() == reflect.Slice {
			if l.m == nil {
				l.m = make(map[string]int)
			}
			index := l.m[p.Name]
			l.m[p.Name] = index + 1
			for v.Len() <= index {
				v.Set(reflect.Append(v, reflect.New(v.Type().Elem()).Elem()))
			}
			structValue = v.Index(index)
			sliceOk = true
		} else {
			structValue = v
		}
		// Strip the "I." from "I.X".
		name = name[len(codec.byIndex[decoder.index].name):]
		codec = decoder.substructCodec
	}

	var slice reflect.Value
	if v.Kind() == reflect.Slice && v.Type().Elem().Kind() != reflect.Uint8 {
		slice = v
		v = reflect.New(v.Type().Elem()).Elem()
	} else if _, ok := prev[p.Name]; ok && !sliceOk {
		// Zero the field back out that was set previously, turns out its a slice and we don't know what to do with it
		v.Set(reflect.Zero(v.Type()))

		return "multiple-valued property requires a slice field type"
	}

	prev[p.Name] = struct{}{}

	pValue := p.Value
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x, ok := pValue.(int64)
		if !ok && pValue != nil {
			return typeMismatchReason(p, v)
		}
		if v.OverflowInt(x) {
			return fmt.Sprintf("value %v overflows struct field of type %v", x, v.Type())
		}
		v.SetInt(x)
	case reflect.Bool:
		x, ok := pValue.(bool)
		if !ok && pValue != nil {
			return typeMismatchReason(p, v)
		}
		v.SetBool(x)
	case reflect.String:
		switch x := pValue.(type) {
		case BlobKey:
			v.SetString(string(x))
		case string:
			v.SetString(x)
		default:
			if pValue != nil {
				return typeMismatchReason(p, v)
			}
		}
	case reflect.Float32, reflect.Float64:
		x, ok := pValue.(float64)
		if !ok && pValue != nil {
			return typeMismatchReason(p, v)
		}
		if v.OverflowFloat(x) {
			return fmt.Sprintf("value %v overflows struct field of type %v", x, v.Type())
		}
		v.SetFloat(x)
	case reflect.Ptr:
		x, ok := pValue.(*Key)
		if !ok && pValue != nil {
			return typeMismatchReason(p, v)
		}
		if _, ok := v.Interface().(*Key); !ok {
			return typeMismatchReason(p, v)
		}
		v.Set(reflect.ValueOf(x))
	case reflect.Struct:
		switch v.Type() {
		case typeOfTime:
			x, ok := pValue.(time.Time)
			if !ok && pValue != nil {
				return typeMismatchReason(p, v)
			}
			v.Set(reflect.ValueOf(x))
		default:
			return typeMismatchReason(p, v)
		}
	case reflect.Slice:
		x, ok := pValue.([]byte)
		if !ok && pValue != nil {
			return typeMismatchReason(p, v)
		}
		if v.Type().Elem().Kind() != reflect.Uint8 {
			return typeMismatchReason(p, v)
		}
		v.SetBytes(x)
	default:
		return typeMismatchReason(p, v)
	}
	if slice.IsValid() {
		slice.Set(reflect.Append(slice, v))
	}
	return ""
}

// loadEntity loads an EntityProto into PropertyLoadSaver or struct pointer.
func loadEntity(dst interface{}, src *pb.Entity) (err error) {
	props := protoToProperties(src)
	if e, ok := dst.(PropertyLoadSaver); ok {
		return e.Load(props)
	}
	return LoadStruct(dst, props)
}

func (s structPLS) Load(props []Property) error {
	var fieldName, reason string
	var l propertyLoader

	prev := make(map[string]struct{})
	for _, p := range props {
		if errStr := l.load(s.codec, s.v, p, prev); errStr != "" {
			// We don't return early, as we try to load as many properties as possible.
			// It is valid to load an entity into a struct that cannot fully represent it.
			// That case returns an error, but the caller is free to ignore it.
			fieldName, reason = p.Name, errStr
		}
	}
	if reason != "" {
		return &ErrFieldMismatch{
			StructType: s.v.Type(),
			FieldName:  fieldName,
			Reason:     reason,
		}
	}
	return nil
}

func protoToProperties(src *pb.Entity) []Property {
	props := src.Property
	out := make([]Property, 0, len(props))
	for {
		var (
			x       *pb.Property
			noIndex bool
		)
		if len(props) > 0 {
			x, props = props[0], props[1:]
			noIndex = !x.GetValue().GetIndexed()
		} else {
			break
		}

		value := propValue(x.Value)

		out = append(out, Property{
			Name:    x.GetName(),
			Value:   value,
			NoIndex: noIndex,
		})

	}
	return out
}

// propValue returns a Go value that combines the raw PropertyValue with a
// meaning. For example, an Int64Value with GD_WHEN becomes a time.Time.
func propValue(v *pb.Value) interface{} {
	//TODO(PSG-Luna): Support EntityValue
	//TODO(PSG-Luna): GeoPoint seems gone from the v1 proto, reimplement it once it's readded
	switch {
	case v.IntegerValue != nil:
		return *v.IntegerValue
	case v.TimestampMicrosecondsValue != nil:
		return fromUnixMicro(*v.TimestampMicrosecondsValue)
	case v.BooleanValue != nil:
		return *v.BooleanValue
	case v.StringValue != nil:
		return *v.StringValue
	case v.BlobValue != nil:
		return []byte(v.BlobValue)
	case v.BlobKeyValue != nil:
		return BlobKey(*v.BlobKeyValue)
	case v.DoubleValue != nil:
		return *v.DoubleValue
	case v.KeyValue != nil:
		return protoToKey(v.KeyValue)
	case v.ListValue != nil:
		// Loop through the listvalues until we find a non nil value and then initialize the slice
		// If no non nil values are found, a []interface{} with all nil values of the length of ListValue will be returned
		var slice reflect.Value
		for i := 0; i < len(v.ListValue); i++ {
			val := propValue(v.ListValue[i])

			if val != nil {
				rVal := reflect.ValueOf(val)
				if !slice.IsValid() {
					slice = reflect.MakeSlice(reflect.SliceOf(rVal.Type()), len(v.ListValue), len(v.ListValue))
				}

				slice.Index(i).Set(rVal)
			}
		}

		if !slice.IsValid() {
			return make([]interface{}, len(v.ListValue))
		} else {
			return slice.Interface()
		}
	}
	return nil
}
