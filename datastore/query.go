// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	"code.google.com/p/goprotobuf/proto"
	pb "google.golang.org/cloud/internal/datastore"
)

type operator int

const (
	lessThan operator = iota
	lessEq
	equal
	greaterEq
	greaterThan
)

var operatorToProto = map[operator]*pb.PropertyFilter_Operator{
	lessThan:    pb.PropertyFilter_LESS_THAN.Enum(),
	lessEq:      pb.PropertyFilter_LESS_THAN_OR_EQUAL.Enum(),
	equal:       pb.PropertyFilter_EQUAL.Enum(),
	greaterEq:   pb.PropertyFilter_GREATER_THAN_OR_EQUAL.Enum(),
	greaterThan: pb.PropertyFilter_GREATER_THAN.Enum(),
}

// filter is a conditional filter on query results.
type filter struct {
	FieldName string
	Op        operator
	Value     interface{}
}

type sortDirection int

const (
	ascending sortDirection = iota
	descending
)

var sortDirectionToProto = map[sortDirection]*pb.PropertyOrder_Direction{
	ascending:  pb.PropertyOrder_ASCENDING.Enum(),
	descending: pb.PropertyOrder_DESCENDING.Enum(),
}

// order is a sort order on query results.
type order struct {
	FieldName string
	Direction sortDirection
}

// NewQuery creates a new Query for a specific entity kind.
//
// An empty kind means to return all entities, including entities created and
// managed by other App Engine features, and is called a kindless query.
// Kindless queries cannot include filters or sort orders on property values.
func NewQuery(kind string) *Query {
	return &Query{
		kind:  kind,
		limit: -1,
	}
}

// Query represents a datastore query.
type Query struct {
	kind       string
	ancestor   *Key
	filter     []filter
	order      []order
	projection []string

	distinct bool
	keysOnly bool
	eventual bool
	limit    int32
	offset   int32
	start    []byte
	end      []byte

	err error
}

func (q *Query) clone() *Query {
	x := *q
	// Copy the contents of the slice-typed fields to a new backing store.
	if len(q.filter) > 0 {
		x.filter = make([]filter, len(q.filter))
		copy(x.filter, q.filter)
	}
	if len(q.order) > 0 {
		x.order = make([]order, len(q.order))
		copy(x.order, q.order)
	}
	return &x
}

// Ancestor returns a derivative query with an ancestor filter.
// The ancestor should not be nil.
func (q *Query) Ancestor(ancestor *Key) *Query {
	q = q.clone()
	if ancestor == nil {
		q.err = errors.New("datastore: nil query ancestor")
		return q
	}
	q.ancestor = ancestor
	return q
}

// EventualConsistency returns a derivative query that returns eventually
// consistent results.
// It only has an effect on ancestor queries.
func (q *Query) EventualConsistency() *Query {
	q = q.clone()
	q.eventual = true
	return q
}

// Filter returns a derivative query with a field-based filter.
// The filterStr argument must be a field name followed by optional space,
// followed by an operator, one of ">", "<", ">=", "<=", or "=".
// Fields are compared against the provided value using the operator.
// Multiple filters are AND'ed together.
func (q *Query) Filter(filterStr string, value interface{}) *Query {
	q = q.clone()
	filterStr = strings.TrimSpace(filterStr)
	if len(filterStr) < 1 {
		q.err = errors.New("datastore: invalid filter: " + filterStr)
		return q
	}
	f := filter{
		FieldName: strings.TrimRight(filterStr, " ><=!"),
		Value:     value,
	}
	switch op := strings.TrimSpace(filterStr[len(f.FieldName):]); op {
	case "<=":
		f.Op = lessEq
	case ">=":
		f.Op = greaterEq
	case "<":
		f.Op = lessThan
	case ">":
		f.Op = greaterThan
	case "=":
		f.Op = equal
	default:
		q.err = fmt.Errorf("datastore: invalid operator %q in filter %q", op, filterStr)
		return q
	}
	q.filter = append(q.filter, f)
	return q
}

// Order returns a derivative query with a field-based sort order. Orders are
// applied in the order they are added. The default order is ascending; to sort
// in descending order prefix the fieldName with a minus sign (-).
func (q *Query) Order(fieldName string) *Query {
	q = q.clone()
	fieldName = strings.TrimSpace(fieldName)
	o := order{
		Direction: ascending,
		FieldName: fieldName,
	}
	if strings.HasPrefix(fieldName, "-") {
		o.Direction = descending
		o.FieldName = strings.TrimSpace(fieldName[1:])
	} else if strings.HasPrefix(fieldName, "+") {
		q.err = fmt.Errorf("datastore: invalid order: %q", fieldName)
		return q
	}
	if len(o.FieldName) == 0 {
		q.err = errors.New("datastore: empty order")
		return q
	}
	q.order = append(q.order, o)
	return q
}

// Project returns a derivative query that yields only the given fields. It
// cannot be used with KeysOnly.
func (q *Query) Project(fieldNames ...string) *Query {
	q = q.clone()
	q.projection = append([]string(nil), fieldNames...)
	return q
}

// Distinct returns a derivative query that yields de-duplicated entities with
// respect to the set of projected fields. It is only used for projection
// queries.
func (q *Query) Distinct() *Query {
	q = q.clone()
	q.distinct = true
	return q
}

// KeysOnly returns a derivative query that yields only keys, not keys and
// entities. It cannot be used with projection queries.
func (q *Query) KeysOnly() *Query {
	q = q.clone()
	q.keysOnly = true
	return q
}

// Limit returns a derivative query that has a limit on the number of results
// returned. A negative value means unlimited.
func (q *Query) Limit(limit int) *Query {
	q = q.clone()
	if limit < math.MinInt32 || limit > math.MaxInt32 {
		q.err = errors.New("datastore: query limit overflow")
		return q
	}
	q.limit = int32(limit)
	return q
}

// Offset returns a derivative query that has an offset of how many keys to
// skip over before returning results. A negative value is invalid.
func (q *Query) Offset(offset int) *Query {
	q = q.clone()
	if offset < 0 {
		q.err = errors.New("datastore: negative query offset")
		return q
	}
	if offset > math.MaxInt32 {
		q.err = errors.New("datastore: query offset overflow")
		return q
	}
	q.offset = int32(offset)
	return q
}

// Start returns a derivative query with the given start point.
func (q *Query) Start(c Cursor) *Query {
	q = q.clone()
	if c.cc == nil {
		q.err = errors.New("datastore: invalid cursor")
		return q
	}
	q.start = c.cc
	return q
}

// End returns a derivative query with the given end point.
func (q *Query) End(c Cursor) *Query {
	q = q.clone()
	if c.cc == nil {
		q.err = errors.New("datastore: invalid cursor")
		return q
	}
	q.end = c.cc
	return q
}

// toProto converts the query to a protocol buffer.
func (q *Query) toProto(req *pb.RunQueryRequest) error {
	dst := pb.Query{}

	if q.kind == "" {
		if len(q.filter) != 0 {
			return errors.New("datastore: kindless query cannot have filters")
		}
		if len(q.order) != 0 {
			return errors.New("datastore: kindless query cannot have sort orders")
		}
	}
	if len(q.projection) != 0 && q.keysOnly {
		return errors.New("datastore: query cannot both project and be keys-only")
	}
	dst.Reset()
	if q.kind != "" {
		dst.Kind = []*pb.KindExpression{&pb.KindExpression{Name: proto.String(q.kind)}}
	}

	if q.projection != nil {
		for _, propertyName := range q.projection {
			dst.Projection = append(dst.Projection, &pb.PropertyExpression{Property: &pb.PropertyReference{Name: proto.String(propertyName)}})
		}

		if q.distinct {
			for _, propertyName := range q.projection {
				dst.GroupBy = append(dst.GroupBy, &pb.PropertyReference{Name: proto.String(propertyName)})
			}
		}
	}
	if q.keysOnly {
		dst.Projection = []*pb.PropertyExpression{&pb.PropertyExpression{Property: &pb.PropertyReference{Name: proto.String("__key__")}}}
	}

	var filters []*pb.Filter
	for _, qf := range q.filter {
		if qf.FieldName == "" {
			return errors.New("datastore: empty query filter field name")
		}

		v, errStr := valueToProto(reflect.ValueOf(qf.Value))
		if errStr != "" {
			return errors.New("datastore: bad query filter value type: " + errStr)
		}

		xf := &pb.PropertyFilter{
			Operator: operatorToProto[qf.Op],
			Property: &pb.PropertyReference{Name: proto.String(qf.FieldName)},
			Value:    v,
		}

		if xf.Operator == nil {
			return errors.New("datastore: unknown query filter operator")
		}

		filters = append(filters, &pb.Filter{PropertyFilter: xf})
	}

	if q.ancestor != nil {
		filters = append(filters, &pb.Filter{
			PropertyFilter: &pb.PropertyFilter{
				Property: &pb.PropertyReference{Name: proto.String("__key__")},
				Operator: pb.PropertyFilter_HAS_ANCESTOR.Enum(),
				Value:    &pb.Value{KeyValue: keyToProto(q.ancestor)},
			}})
	}

	if len(filters) == 1 {
		dst.Filter = filters[0]
	} else if len(filters) > 1 {
		dst.Filter = &pb.Filter{CompositeFilter: &pb.CompositeFilter{
			Operator: pb.CompositeFilter_AND.Enum(),
			Filter:   filters,
		}}
	}

	for _, qo := range q.order {
		if qo.FieldName == "" {
			return errors.New("datastore: empty query order field name")
		}
		xo := &pb.PropertyOrder{
			Property:  &pb.PropertyReference{Name: proto.String(qo.FieldName)},
			Direction: sortDirectionToProto[qo.Direction],
		}
		if xo.Direction == nil {
			return errors.New("datastore: unknown query order direction")
		}
		dst.Order = append(dst.Order, xo)
	}
	if q.limit >= 0 {
		dst.Limit = proto.Int32(q.limit)
	}
	if q.offset != 0 {
		dst.Offset = proto.Int32(q.offset)
	}
	dst.StartCursor = q.start
	dst.EndCursor = q.end

	req.Query = &dst
	return nil
}

// Count returns the number of results for the query.
func (q *Query) Count(c Client) (int, error) {
	// Check that the query is well-formed.
	if q.err != nil {
		return 0, q.err
	}

	// Run a copy of the query, with keysOnly true (if we're not a projection,
	// since the two are incompatible), and an adjusted offset. We also set the
	// limit to zero, as we don't want any actual entity data, just the number
	// of skipped results.
	newQ := q.clone()
	newQ.keysOnly = len(newQ.projection) == 0
	newQ.limit = 0
	if q.limit < 0 {
		// If the original query was unlimited, set the new query's offset to maximum.
		newQ.offset = math.MaxInt32
	} else {
		newQ.offset = q.offset + q.limit
		if newQ.offset < 0 {
			// Do the best we can, in the presence of overflow.
			newQ.offset = math.MaxInt32
		}
	}
	req := &pb.RunQueryRequest{}
	if err := newQ.toProto(req); err != nil {
		return 0, err
	}
	res := &pb.RunQueryResponse{}
	if err := c.Call("RunQuery", req, res); err != nil {
		return 0, err
	}

	// n is the count we will return. For example, suppose that our original
	// query had an offset of 4 and a limit of 2008: the count will be 2008,
	// provided that there are at least 2012 matching entities. However, the
	// RPCs will only skip 1000 results at a time. The RPC sequence is:
	//   call RunQuery with (offset, limit) = (2012, 0)  // 2012 == newQ.offset
	//   response has (skippedResults, moreResults) = (1000, true)
	//   n += 1000  // n == 1000
	//   call Next     with (offset, limit) = (1012, 0)  // 1012 == newQ.offset - n
	//   response has (skippedResults, moreResults) = (1000, true)
	//   n += 1000  // n == 2000
	//   call Next     with (offset, limit) = (12, 0)    // 12 == newQ.offset - n
	//   response has (skippedResults, moreResults) = (12, false)
	//   n += 12    // n == 2012
	//   // exit the loop
	//   n -= 4     // n == 2008
	var n int32
	b := res.GetBatch()
	for {
		// The QueryResult should have no actual entity data, just skipped results.
		if len(b.EntityResult) != 0 {
			return 0, errors.New("datastore: internal error: Count request returned too much data")
		}
		n += b.GetSkippedResults()
		if b.GetMoreResults() != pb.QueryResultBatch_NOT_FINISHED {
			break
		}
		var err error
		if err = callNext(c, req, res, newQ.offset-n, 0); err != nil {
			return 0, err
		}
	}
	n -= q.offset
	if n < 0 {
		// If the offset was greater than the number of matching entities,
		// return 0 instead of negative.
		n = 0
	}
	return int(n), nil
}

// callNext issues a datastore_v3/Next RPC to advance a cursor, such as that
// returned by a query with more results.
func callNext(c Client, req *pb.RunQueryRequest, res *pb.RunQueryResponse, offset, limit int32) error {
	if res.GetBatch().EndCursor == nil {
		return errors.New("datastore: internal error: server did not return a cursor")
	}

	req.GetQuery().StartCursor = res.GetBatch().GetEndCursor()

	if limit >= 0 {
		req.GetQuery().Limit = proto.Int32(limit)
	}
	if offset != 0 {
		req.GetQuery().Offset = proto.Int32(offset)
	}

	res.Reset()
	return c.Call("RunQuery", req, res)
}

// GetAll runs the query in the given context and returns all keys that match
// that query, as well as appending the values to dst.
//
// dst must have type *[]S or *[]*S or *[]P, for some struct type S or some non-
// interface, non-pointer type P such that P or *P implements PropertyLoadSaver.
//
// As a special case, *PropertyList is an invalid type for dst, even though a
// PropertyList is a slice of structs. It is treated as invalid to avoid being
// mistakenly passed when *[]PropertyList was intended.
//
// The keys returned by GetAll will be in a 1-1 correspondence with the entities
// added to dst.
//
// If q is a ``keys-only'' query, GetAll ignores dst and only returns the keys.
func (q *Query) GetAll(c Client, dst interface{}) ([]*Key, error) {
	var (
		dv               reflect.Value
		mat              multiArgType
		elemType         reflect.Type
		errFieldMismatch error
	)
	if !q.keysOnly {
		dv = reflect.ValueOf(dst)
		if dv.Kind() != reflect.Ptr || dv.IsNil() {
			return nil, ErrInvalidEntityType
		}
		dv = dv.Elem()
		mat, elemType = checkMultiArg(dv)
		if mat == multiArgTypeInvalid || mat == multiArgTypeInterface {
			return nil, ErrInvalidEntityType
		}
	}

	var keys []*Key
	for t := q.Run(c); ; {
		k, e, err := t.next()
		if err == Done {
			break
		}
		if err != nil {
			return keys, err
		}
		if !q.keysOnly {
			ev := reflect.New(elemType)
			if elemType.Kind() == reflect.Map {
				// This is a special case. The zero values of a map type are
				// not immediately useful; they have to be make'd.
				//
				// Funcs and channels are similar, in that a zero value is not useful,
				// but even a freshly make'd channel isn't useful: there's no fixed
				// channel buffer size that is always going to be large enough, and
				// there's no goroutine to drain the other end. Theoretically, these
				// types could be supported, for example by sniffing for a constructor
				// method or requiring prior registration, but for now it's not a
				// frequent enough concern to be worth it. Programmers can work around
				// it by explicitly using Iterator.Next instead of the Query.GetAll
				// convenience method.
				x := reflect.MakeMap(elemType)
				ev.Elem().Set(x)
			}
			if err = loadEntity(ev.Interface(), e); err != nil {
				if _, ok := err.(*ErrFieldMismatch); ok {
					// We continue loading entities even in the face of field mismatch errors.
					// If we encounter any other error, that other error is returned. Otherwise,
					// an ErrFieldMismatch is returned.
					errFieldMismatch = err
				} else {
					return keys, err
				}
			}
			if mat != multiArgTypeStructPtr {
				ev = ev.Elem()
			}
			dv.Set(reflect.Append(dv, ev))
		}
		keys = append(keys, k)
	}
	return keys, errFieldMismatch
}

// Run runs the query in the given context.
func (q *Query) Run(c Client) *Iterator {
	if q.err != nil {
		return &Iterator{err: q.err}
	}
	t := &Iterator{
		c:      c,
		limit:  q.limit,
		q:      q,
		prevCC: q.start,
	}
	t.req.Reset()
	if err := q.toProto(&t.req); err != nil {
		t.err = err
		return t
	}
	if err := c.Call("RunQuery", &t.req, &t.res); err != nil {
		t.err = err
		return t
	}
	b := t.res.GetBatch()
	offset := q.offset - b.GetSkippedResults()
	for offset > 0 && b.GetMoreResults() == pb.QueryResultBatch_NOT_FINISHED {
		t.prevCC = b.GetEndCursor()
		var err error
		if err = callNext(t.c, &t.req, &t.res, offset, t.limit); err != nil {
			t.err = err
			break
		}
		skip := b.GetSkippedResults()
		if skip < 0 {
			t.err = errors.New("datastore: internal error: negative number of skipped_results")
			break
		}
		offset -= skip
	}
	if offset < 0 {
		t.err = errors.New("datastore: internal error: query offset was overshot")
	}
	return t
}

// Iterator is the result of running a query.
type Iterator struct {
	c   Client
	err error
	// req is the request we sent previously, we need to keep track of it to resend it
	req pb.RunQueryRequest
	// res is the result of the most recent RunQuery or Next API call.
	res pb.RunQueryResponse
	// i is how many elements of res.Result we have iterated over.
	i int
	// limit is the limit on the number of results this iterator should return.
	// A negative value means unlimited.
	limit int32
	// q is the original query which yielded this iterator.
	q *Query
	// prevCC is the compiled cursor that marks the end of the previous batch
	// of results.
	prevCC []byte
}

// Done is returned when a query iteration has completed.
var Done = errors.New("datastore: query has no more results")

// Next returns the key of the next result. When there are no more results,
// Done is returned as the error.
//
// If the query is not keys only and dst is non-nil, it also loads the entity
// stored for that key into the struct pointer or PropertyLoadSaver dst, with
// the same semantics and possible errors as for the Get function.
func (t *Iterator) Next(dst interface{}) (*Key, error) {
	k, e, err := t.next()
	if err != nil {
		return nil, err
	}
	if dst != nil && !t.q.keysOnly {
		err = loadEntity(dst, e)
	}
	return k, err
}

func (t *Iterator) next() (*Key, *pb.Entity, error) {
	if t.err != nil {
		return nil, nil, t.err
	}

	// Issue datastore_v3/Next RPCs as necessary.
	b := t.res.GetBatch()
	for t.i == len(b.EntityResult) {
		if b.GetMoreResults() != pb.QueryResultBatch_NOT_FINISHED {
			t.err = Done
			return nil, nil, t.err
		}
		t.prevCC = b.GetEndCursor()
		if err := callNext(t.c, &t.req, &t.res, 0, t.limit); err != nil {
			t.err = err
			return nil, nil, t.err
		}
		if b.GetSkippedResults() != 0 {
			t.err = errors.New("datastore: internal error: iterator has skipped results")
			return nil, nil, t.err
		}
		t.i = 0
		if t.limit >= 0 {
			t.limit -= int32(len(b.EntityResult))
			if t.limit < 0 {
				t.err = errors.New("datastore: internal error: query returned more results than the limit")
				return nil, nil, t.err
			}
		}
	}

	// Extract the key from the t.i'th element of t.res.Result.
	e := b.EntityResult[t.i]
	t.i++
	if e.Entity.Key == nil {
		return nil, nil, errors.New("datastore: internal error: server did not return a key")
	}
	k := protoToKey(e.Entity.Key)
	if k.Incomplete() {
		return nil, nil, errors.New("datastore: internal error: server returned an invalid key")
	}
	return k, e.Entity, nil
}

// Cursor returns a cursor for the iterator's current location.
func (t *Iterator) Cursor() (Cursor, error) {
	if t.err != nil && t.err != Done {
		return Cursor{}, t.err
	}
	// If we are at either end of the current batch of results,
	// return the compiled cursor at that end.
	b := t.res.Batch
	skipped := b.GetSkippedResults()
	if t.i == 0 && skipped == 0 {
		if t.prevCC == nil {
			// A nil pointer (of type *pb.CompiledCursor) means no constraint:
			// passing it as the end cursor of a new query means unlimited results
			// (glossing over the integer limit parameter for now).
			// A non-nil pointer to an empty pb.CompiledCursor means the start:
			// passing it as the end cursor of a new query means 0 results.
			// If prevCC was nil, then the original query had no start cursor, but
			// Iterator.Cursor should return "the start" instead of unlimited.
			return Cursor{}, nil
		}
		return Cursor{t.prevCC}, nil
	}
	if t.i == len(b.EntityResult) {
		return Cursor{b.EndCursor}, nil
	}
	// Otherwise, re-run the query offset to this iterator's position, starting from
	// the most recent compiled cursor. This is done on a best-effort basis, as it
	// is racy; if a concurrent process has added or removed entities, then the
	// cursor returned may be inconsistent.
	q := t.q.clone()
	q.start = t.prevCC
	q.offset = skipped + int32(t.i)
	q.limit = 0
	q.keysOnly = len(q.projection) == 0
	t1 := q.Run(t.c)
	_, _, err := t1.next()
	if err != Done {
		if err == nil {
			err = fmt.Errorf("datastore: internal error: zero-limit query did not have zero results")
		}
		return Cursor{}, err
	}
	return Cursor{t1.res.Batch.EndCursor}, nil
}

// Cursor is an iterator's position. It can be converted to and from an opaque
// string. A cursor can be used from different HTTP requests, but only with a
// query with the same kind, ancestor, filter and order constraints.
type Cursor struct {
	cc []byte
}

// String returns a base-64 string representation of a cursor.
func (c Cursor) String() string {
	if c.cc == nil {
		return ""
	}

	return strings.TrimRight(base64.URLEncoding.EncodeToString(c.cc), "=")
}

// Decode decodes a cursor from its base-64 string representation.
func DecodeCursor(s string) (Cursor, error) {
	if s == "" {
		return Cursor{}, nil
	}
	if n := len(s) % 4; n != 0 {
		s += strings.Repeat("=", 4-n)
	}
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, err
	}
	return Cursor{b}, nil
}
