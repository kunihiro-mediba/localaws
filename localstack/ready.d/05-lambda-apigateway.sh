#!/bin/bash

awslocal lambda create-function \
    --function-name api \
    --runtime nodejs22.x \
    --handler index.handler \
    --role arn:aws:iam::000000000000:role/lambda-execution-role \
    --zip-file fileb:///app/init/lambda/hello.zip

awslocal apigateway create-rest-api \
    --name 'Lambda backend API example' \
    --tags '{"_custom_id_":"apiid"}' \
    | tee create-api.json

export apiid=$(jq -r .id create-api.json)
export parentid=$(jq -r .rootResourceId create-api.json)
echo "apiid=${apiid} parentid=${parentid}"

awslocal apigateway create-resource \
    --rest-api-id ${apiid} \
    --parent-id ${parentid} \
    --path-part "{proxy+}" \
    | tee create-resource.json

export resourceid=$(jq -r .id create-resource.json)
echo "resourceid=${resourceid}"

awslocal apigateway put-method \
    --rest-api-id ${apiid} \
    --resource-id ${resourceid} \
    --http-method ANY \
    --authorization-type "NONE" \
    | tee put-method.json

awslocal apigateway put-integration \
    --rest-api-id ${apiid} \
    --resource-id ${resourceid} \
    --http-method ANY \
    --type AWS_PROXY \
    --integration-http-method POST \
    --uri arn:aws:apigateway:ap-northeast-1:lambda:path/2015-03-31/functions/arn:aws:lambda:ap-northeast-1:000000000000:function:api \
    --credentials arn:aws:iam::000000000000:role/apigAwsProxyRole \
    | tee put-integration.json

awslocal apigateway create-deployment \
    --rest-api-id ${apiid} \
    --stage-name dev

# open http://apiid.execute-api.localhost.localstack.cloud:4566/dev/hello
