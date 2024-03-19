#!/bin/bash

awslocal s3 mb s3://my-bucket

awslocal s3 ls

awslocal s3 ls s3://my-bucket
