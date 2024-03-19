#!/bin/bash

PARAMETER_NAME="parameter-name-test"

awslocal ssm put-parameter \
  --type "SecureString" \
  --name $PARAMETER_NAME \
  --value "Hello,SSM!"

awslocal ssm get-parameter \
  --with-decryption \
  --name $PARAMETER_NAME
