package main

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

const (
	KMSKey    = "alias/kmskey/test"
	PlainText = "Hello,KMS!"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig %v", err)
		return
	}

	client := kms.NewFromConfig(cfg)

	// Encrypt
	cipherTextBase64 := ""
	if out, err := client.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:     aws.String(KMSKey),
		Plaintext: []byte(PlainText),
	}); err != nil {
		log.Fatalf("kms.Client.Encrypt %v", err)
		return
	} else {
		cipherTextBase64 = base64.StdEncoding.EncodeToString(out.CiphertextBlob)
		log.Printf("cipertext: %s", cipherTextBase64)
	}

	// Decrypt
	cipherText, _ := base64.StdEncoding.DecodeString(cipherTextBase64)
	if out, err := client.Decrypt(context.TODO(), &kms.DecryptInput{
		CiphertextBlob: cipherText,
	}); err != nil {
		log.Fatalf("kms.Client.Decrypt %v", err)
		return
	} else {
		log.Printf("plaintext: %s", string(out.Plaintext))
	}
}
