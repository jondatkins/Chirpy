#!/usr/bin/env bash

curl -w "\n" --request POST http://localhost:8080/admin/reset
curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users
curl -w "\n" --data '{ "email": "walt@breakingbad.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users

RESPONSE=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"walt@breakingbad.com","password":"password1"}')

USER_ID=$(echo "$RESPONSE" | jq -r '.id')
JWT_TOKEN=$(echo "$RESPONSE" | jq -r '.token')
# echo $JWT_TOKEN
FORM_DATA=$(jq -n --arg uid "$USER_ID" --arg body "My name is Walt!!!" '{user_id: $uid, body: $body}')

curl -s -X POST http://localhost:8080/api/chirps \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d "$FORM_DATA"

RESPONSE=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"saul@bettercall.com","password":"password1"}')

USER_ID=$(echo "$RESPONSE" | jq -r '.id')
JWT_TOKEN=$(echo "$RESPONSE" | jq -r '.token')
# echo $JWT_TOKEN
FORM_DATA=$(jq -n --arg uid "$USER_ID" --arg body "My name is Saul!!!" '{user_id: $uid, body: $body}')

curl -s -X POST http://localhost:8080/api/chirps \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d "$FORM_DATA"

# curl -s -X DELETE http://localhost:8080/api/chirps/1f241f37-29af-4199-b602-6ea9b5e2e9a6

DELETE_RESP=$(curl -s -X DELETE http://localhost:8080/api/chirps/1f241f37-29af-4199-b602-6ea9b5e2e9a6 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN") \
  echo "FOO"
echo $DELETE_RESP
