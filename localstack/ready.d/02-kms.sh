#!/bin/bash

KeyName="alias/kmskey/test"

cd /app/init/keys

openssl enc -d -base64 -A -in PlaintextKeyMaterial.b64 -out PlaintextKeyMaterial.bin

awslocal kms create-key \
  --origin EXTERNAL \
  | tee create-key-output.json
KeyId=`jq -r '.KeyMetadata.KeyId' create-key-output.json`

awslocal kms get-parameters-for-import \
  --wrapping-algorithm RSAES_OAEP_SHA_1 \
  --wrapping-key-spec RSA_2048 \
  --key-id ${KeyId} \
  | tee get-parameters-for-import-output.json

jq -r '.PublicKey' get-parameters-for-import-output.json > PublicKey.b64
openssl enc -d -base64 -A -in PublicKey.b64 -out PublicKey.bin

jq -r '.ImportToken' get-parameters-for-import-output.json > ImportToken.b64
openssl enc -d -base64 -A -in ImportToken.b64 -out ImportToken.bin

openssl rsautl -encrypt \
  -in PlaintextKeyMaterial.bin \
  -oaep \
  -inkey PublicKey.bin \
  -keyform DER \
  -pubin \
  -out EncryptedKeyMaterial.bin

awslocal kms import-key-material \
  --encrypted-key-material fileb://EncryptedKeyMaterial.bin \
  --import-token fileb://ImportToken.bin \
  --expiration-model KEY_MATERIAL_DOES_NOT_EXPIRE \
  --key-id ${KeyId}

awslocal kms create-alias \
  --target-key-id ${KeyId} \
  --alias-name ${KeyName}

awslocal kms describe-key \
  --key-id ${KeyName}
