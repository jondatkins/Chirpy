#!/usr/bin/env bash
curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/login | jq '.["token"]'
# curl http://localhost:8080/api/users
