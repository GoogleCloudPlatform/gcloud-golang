module cloud.google.com/go/internal/generated

go 1.16

replace cloud.google.com/go/bigquery => ../../../bigquery

replace cloud.google.com/go/bigtable => ../../../bigtable

replace cloud.google.com/go/datastore => ../../../datastore

replace cloud.google.com/go/firestore => ../../../firestore

replace cloud.google.com/go => ../../..

replace cloud.google.com/go/logging => ../../../logging

replace cloud.google.com/go/pubsub => ../../../pubsub

replace cloud.google.com/go/pubsublite => ../../../pubsublite

replace cloud.google.com/go/spanner => ../../../spanner

replace cloud.google.com/go/storage => ../../../storage

require (
	cloud.google.com/go v0.81.0
	cloud.google.com/go/bigquery v0.81.0
	cloud.google.com/go/datastore v0.81.0
	cloud.google.com/go/firestore v0.81.0
	cloud.google.com/go/logging v0.81.0
	cloud.google.com/go/pubsub v1.9.1
	cloud.google.com/go/pubsublite v0.81.0
	cloud.google.com/go/spanner v0.81.0
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384
)
