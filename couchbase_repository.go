package main

import (
	couchbase "github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"strings"
)

type CouchbaseRepository struct {
	cbClient *couchbase.Cluster
}

type ICouchbaseRepository interface {
	FilterByRate(int) ([]Review, error)
	Upsert(Review) (*Review, error)
	GetById(string) (*Review, error)
}

func NewCouchbaseRepositoryAdaptor(cbClient *couchbase.Cluster) ICouchbaseRepository {
	return &CouchbaseRepository{
		cbClient: cbClient,
	}
}

func (c *CouchbaseRepository) FilterByRate(rate int) ([]Review, error) {
	params := make(map[string]interface{}, 1)
	params["rate"] = rate

	query := "select `r`.* from reviews r where rate = $rate order by lastModifiedDate limit 20;"
	rows, err := c.cbClient.Query(query, &couchbase.QueryOptions{NamedParameters: params})

	if err != nil {
		log.Error("An error occurred while retrieving document from couchbase")
		return nil, err
	}

	var reviews []Review
	if rows == nil {
		return reviews, nil
	}

	if err != nil && !strings.Contains(err.Error(), "document not found") {
		log.Errorf("Error while getting blacklist users response from couchbase {}", err)
	}
	if rows != nil {

		for rows.Next() {
			var row Review

			err = rows.Row(&row)
			if err != nil {
				log.Errorf("Error while getting blacklist user row from couchbase {}", err)
			}
			reviews = append(reviews, row)
		}
	}

	return reviews, nil
}

func (c *CouchbaseRepository) GetById(id string) (review *Review, errOut error) {
	getResult, err := c.cbClient.Bucket("reviews").DefaultCollection().Get(id, &couchbase.GetOptions{})

	if err != nil {
		log.Error("An error occurred while retrieving document from couchbase with id: %s", id)
		return nil, err
	}

	if getResult == nil {
		return nil, nil
	}

	var r *Review
	err = getResult.Content(&r)
	if err != nil {
		log.Error("An error occurred while converting review with id: %s", id)
		return nil, err
	}
	return r, nil
}

func (c *CouchbaseRepository) Upsert(review Review) (*Review, error) {
	id := uuid.New().String()
	_, err := c.cbClient.Bucket("reviews").DefaultCollection().Upsert(id, review, &couchbase.UpsertOptions{})

	if err != nil {
		log.Error("An error occurred while inserting document to couchbase with id: ", id)
		return nil, err
	}

	return &review, nil
}
