#!/bin/bash

awslocal lambda create-function \
    --function-name batch \
    --runtime nodejs22.x \
    --handler index.handler \
    --role arn:aws:iam::000000000000:role/lambda-execution-role \
    --zip-file fileb:///app/init/lambda/hello.zip

awslocal events put-rule \
    --name batch-lambda-events-rule \
    --schedule-expression 'rate(5 minutes)'

awslocal lambda add-permission \
    --function-name batch \
    --statement-id batch-lambda-event \
    --action 'lambda:InvokeFunction' \
    --principal events.amazonaws.com \
    --source-arn arn:aws:events:ap-northeast-1:000000000000:rule/batch-lambda-events-rule

awslocal events put-targets \
    --rule batch-lambda-events-rule \
    --targets '[
        {
            "Arn": "arn:aws:lambda:ap-northeast-1:000000000000:function:batch",
            "Id": "batch-lambda-event"
        }
    ]'

awslocal events list-rules
