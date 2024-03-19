package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TableName = "test"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig %v", err)
		return
	}

	client := dynamodb.NewFromConfig(cfg)

	// List table name
	if out, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{}); err != nil {
		log.Fatalf("dynamodb.Client.ListTables %v", err)
		return
	} else {
		for _, name := range out.TableNames {
			log.Println(name)
		}
	}

	// Insert item
	if _, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: "1"},
			"name": &types.AttributeValueMemberS{Value: "dynamo"},
		},
	}); err != nil {
		log.Fatalf("dynamodb.Client.PutItem %v", err)
	}

	// Get item
	if out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
	}); err != nil {
		log.Fatalf("dynamodb.Client.GetItem %v", err)
	} else {
		log.Println(out.Item["name"].(*types.AttributeValueMemberS).Value)
	}
}
