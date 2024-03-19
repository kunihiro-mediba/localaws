package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

const (
	SecretID = "secret/key/test"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig %v", err)
		return
	}

	client := secretsmanager.NewFromConfig(cfg)

	if out, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretID),
	}); err != nil {
		log.Fatalf("secretsmanager.Client.GetSecretValue %v", err)
		return
	} else {
		log.Println(*out.SecretString)
	}
}
