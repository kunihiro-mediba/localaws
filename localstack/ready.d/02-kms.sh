#!/bin/bash

KeyId="00000000-0000-0000-0000-000000000001"
KeyMaterial="AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
AliasName="alias/kmskey/test"

awslocal kms create-key \
  --tags "[
    {\"TagKey\":\"_custom_id_\", \"TagValue\":\"${KeyId}\"},
    {\"TagKey\":\"_custom_key_material_\", \"TagValue\":\"${KeyMaterial}\"}
  ]"

awslocal kms create-alias \
  --target-key-id ${KeyId} \
  --alias-name ${AliasName}

awslocal kms describe-key \
  --key-id ${AliasName}
