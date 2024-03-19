package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

const (
	ParameterName = "parameter-name-test"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig %v", err)
		return
	}

	client := ssm.NewFromConfig(cfg)

	// Get parameter
	if out, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(ParameterName),
		WithDecryption: aws.Bool(true),
	}); err != nil {
		log.Fatalf("ssm.Client.GetParameter %v", err)
		return
	} else {
		log.Println(*out.Parameter.Value)
	}
}
