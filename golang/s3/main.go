package main

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	BucketName = "my-bucket"
	ObjectKey  = "filename.txt"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig %v", err)
		return
	}

	client := s3.NewFromConfig(cfg, func(opts *s3.Options) {
		opts.UsePathStyle = true
	})

	// List bucket name
	if out, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{}); err != nil {
		log.Fatalf("s3.Client.ListBuckets %v", err)
		return
	} else {
		for _, bucket := range out.Buckets {
			log.Printf("bucket: %s\n", *bucket.Name)
		}
	}

	// Put object
	if _, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(ObjectKey),
		Body:   strings.NewReader("Hello,World!\n"),
	}); err != nil {
		log.Fatalf("s3.Client.PutObject %v", err)
		return
	}

	// Get object
	if out, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(ObjectKey),
	}); err != nil {
		log.Fatalf("s3.Client.GetObject %v", err)
		return
	} else {
		if buf, err := io.ReadAll(out.Body); err != nil {
			log.Fatalf("io.ReadAll %v", err)
			return
		} else {
			log.Printf("body: %s", string(buf))
		}
	}
}
