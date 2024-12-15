#!/bin/bash

for i in {1..4}; do
    curl -is -w "Request $i: %{http_code}\n" -o /dev/null "http://localhost:8080/api/v1/healthz"
done

echo "wait for block duration: 5s"
sleep 5

curl -is -w "status: %{http_code}\n" -o /dev/null http://localhost:8080/api/v1/healthz