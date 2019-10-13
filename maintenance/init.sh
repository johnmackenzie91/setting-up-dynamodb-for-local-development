#!/bin/sh
AWS_ACCESS_KEY_ID=AWS_ACCESS_KEY_ID \
AWS_SECRET_ACCESS_KEY=AWS_SECRET_ACCESS_KEY \
aws dynamodb delete-table \
	--table-name football-results \
	--endpoint-url http://0.0.0.0:8000 \
	--region us-west-2

AWS_ACCESS_KEY_ID=AWS_ACCESS_KEY_ID \
AWS_SECRET_ACCESS_KEY=AWS_SECRET_ACCESS_KEY \
aws dynamodb create-table \
	--cli-input-json "$(cat 01-create-table.json)" \
	--endpoint-url http://0.0.0.0:8000 \
	--region us-west-2

AWS_ACCESS_KEY_ID=AWS_ACCESS_KEY_ID \
AWS_SECRET_ACCESS_KEY=AWS_SECRET_ACCESS_KEY \
aws dynamodb batch-write-item \
	--request-items "$(cat 02-load-data.json)" \
	--endpoint-url http://0.0.0.0:8000 \
	--region us-west-2

AWS_ACCESS_KEY_ID=AWS_ACCESS_KEY_ID \
AWS_SECRET_ACCESS_KEY=AWS_SECRET_ACCESS_KEY \
aws dynamodb query \
	--table-name football-results \
	--key-condition-expression "partition_key = :league and begins_with(sort_key, :prefix)" \
	--expression-attribute-values "{\":league\": {\"S\": \"english-premier-league\"}, \":prefix\": {\"S\":\"info\"}}" \
	--endpoint-url=http://0.0.0.0:8000 --region us-west-2