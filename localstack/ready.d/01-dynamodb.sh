#!/bin/bash

awslocal dynamodb create-table \
  --cli-input-json file:///app/init/json/dynamodb-table.json

awslocal dynamodb list-tables
