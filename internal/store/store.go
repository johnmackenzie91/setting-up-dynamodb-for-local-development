package store

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"no_vcs/me/dynamo-db-example/internal/models"
)

type LeagueInfoGetter interface {
	GetLeagueInfo(name string) (*models.League, error)
}

type store struct {
	db *dynamodb.DynamoDB
	storeName *string
}

func NewLeagueInfoGetter(name string, db *dynamodb.DynamoDB) LeagueInfoGetter {
	return store{
		db: db,
		storeName: aws.String(name),
	}
}

func (s store) GetLeagueInfo(id string) (*models.League, error) {

	params := &dynamodb.QueryInput{
		TableName:              s.storeName,
		KeyConditionExpression: aws.String("partition_key = :s and begins_with(sort_key, :prefix)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(id),
			},
			":prefix": {
				S: aws.String("info"),
			},
		},
	}

	resp, err := s.db.Query(params)
	if err != nil {
		return nil, err
	}

	var res []models.League
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &res)

	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}