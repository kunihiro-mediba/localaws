#!/bin/bash

SECRET_NAME="secret/key/test"

awslocal secretsmanager create-secret \
  --name $SECRET_NAME \
  --secret-string 'Hello,SecretsManager!'

awslocal secretsmanager get-secret-value \
  --secret-id $SECRET_NAME
