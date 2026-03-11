#!/usr/bin/env bash
curl -w "\n" --request POST http://localhost:8080/admin/reset
curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users | jq '.["id"]'
curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/login | jq '.["token"]'
