package models

import (
	"encoding/json"
	"errors"
	"strings"
)

type League struct {
	PartitionKey string `dynamodbav:"partition_key"`
	SortKey string `dynamodbav:"sort_key"`
	Name string  `dynamodbav:"name"`
	Country string  `dynamodbav:"country"`
}

func (l League) MarshalJSON () ([]byte, error) {
	id, err := l.id()

	if err != nil {
		return nil, err
	}
	return json.Marshal(struct{
		ID string
		Name string
		Country string
	}{
		ID: id,
		Name: l.Name,
		Country: l.Country,
	})
}

func (l League) id () (string, error) {
	parts := strings.Split(l.SortKey, "-")

	if parts[1] == "" {
		return "", errors.New("unable to marshal row, unable to split sort_key")
	}
	return parts[1], nil
}