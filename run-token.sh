#!/bin/bash

API_KEY="my-token"

for i in {1..6}; do
    curl -is -w "Request $i: %{http_code}\n" -o /dev/null -H "API_KEY: $API_KEY" http://localhost:8080/api/v1/healthz
done

echo "wait for block duration: 5s"
sleep 5

curl -is -w "status: %{http_code}\n" -o /dev/null -H "API_KEY: $API_KEY" http://localhost:8080/api/v1/healthz