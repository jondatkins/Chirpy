#!/usr/bin/env bash
curl -w "\n" --request POST http://localhost:8080/admin/reset
saulsId=$(curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users | jq '.["id"]')
waltsId=$(curl -w "\n" --data '{ "email": "walt@breakingbad.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users | jq '.["id"]')
saulsToken=$(curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/login | jq '.["token"]')
echo "sauls id"
echo $saulsId
echo "walts id"
echo $waltsId
echo "sauls token"
echo $saulsToken
curl -w "\n" -H "Authorization: Bearer $saulsToken" --data '{ "body": "FOO BAR", "user_id": "'"$saulsId"'"}' --header 'Content-Type: application/json' http://localhost:8080/api/chirps

# curl -X POST -H "Content-Type: application/json" -d '{"id": "'"$id"'", "jwt": "'"$jwt"'"}' http://example.com/api
