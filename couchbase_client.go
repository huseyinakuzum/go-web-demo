package main

import (
	couchbase "github.com/couchbase/gocb/v2"
	"github.com/labstack/gommon/log"
	"time"
)

func NewCouchbaseClient(couchbaseConfig CouchbaseConfig) *couchbase.Cluster {
	cbClient, err := couchbase.Connect(
		couchbaseConfig.Addresses,
		couchbase.ClusterOptions{
			Username: couchbaseConfig.Username,
			Password: couchbaseConfig.Password,
		})
	if err != nil {
		log.Errorf("Error while connecting to cb {}", err)
		panic(err)
	}

	bucket := cbClient.Bucket("reviews")
	err = bucket.WaitUntilReady(3*time.Second, nil)

	return cbClient
}
