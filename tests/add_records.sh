#!/bin/bash

# Set the API endpoint
API_ENDPOINT='localhost:8080/api/v1/create-record'

# Set the authorization token
AUTH_TOKEN='Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE5OTQ5NDksInVzZXJfaWQiOjF9.ZTFQD9S69OXUIvIk1t9liHm1sf2R-E4yAvOAat5rGzU'

# Check if a JSON file name is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <json_file>"
  exit 1
fi

cat "$1" | jq -c '.[]' | while IFS= read -r record; do
    # Send a request for each record
    curl --location "$API_ENDPOINT" \
         --header 'Content-Type: application/json' \
         --header "Authorization: $AUTH_TOKEN" \
         --data "$record"
done