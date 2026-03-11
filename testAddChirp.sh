#!/usr/bin/env bash
# curl -w "\n" --request POST http://localhost:8080/admin/reset
# user_id=$(curl -w "\n" --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/users | jq '.["id"]')
token=$(curl --data '{ "email": "saul@bettercall.com", "password": "password1" }' --header 'Content-Type: application/json' http://localhost:8080/api/login | jq '.["token"]')
echo $user_id
echo $token
# curl -w "\n" --data '{ "body": "FOO BAR" }' --header 'Content-Type: application/json' http://localhost:8080/api/chirps
# curl -w "\n" --data "{ \"body\": \"FOO BAR\", \"user_id\": \"$user_id\", \"jwt\":\"$token\"}" --header 'Content-Type: application/json' http://localhost:8080/api/chirps

curl -w "\n" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiNDA2ZTA0NTctMTgyZS00NjNkLWJkZWUtMmNiZjEyMTBkMmIyIiwiZXhwIjoxNzczMjQ3Mzg3LCJpYXQiOjE3NzMyNDM3ODd9.ju_-fcQ5sV80cQY4_6iXil-53xz2mc9anCq25UHMY6U" --data '{ "body": "FOO BAR", "user_id": "406e0457-182e-463d-bdee-2cbf1210d2b2"}' --header 'Content-Type: application/json' http://localhost:8080/api/chirps

# eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHktYWNjZXNzIiwic3ViIjoiNDA2ZTA0NTctMTgyZS00NjNkLWJkZWUtMmNiZjEyMTBkMmIyIiwiZXhwIjoxNzczMjQ3Mzg3LCJpYXQiOjE3NzMyNDM3ODd9.ju_-fcQ5sV80cQY4_6iXil-53xz2mc9anCq25UHMY6U
# curl -w "\n" --data '{ "body": "FOO BAR", "user_id": "'"$user_id"'", "jwt": "'"$token"'" }' --header 'Content-Type: application/json' http://localhost:8080/api/chirps
# NEXT_JOB_COMMANDS='["sleep", "200s"]'
# formdata=$(jq -c -n --argjson user_d "$NEXT_JOB_COMMANDS" '$ARGS.named')
# echo $formdata

# curl -H "Authorization: Bearer abcdef123456" https://api.example.com/data
# add_chirp_resp=$(curl -w "\n" -H "Authorization: Bearer 123" --data "$(jq -n --arg token "$token" --arg user_id "$user_id" '{body: "FOO BAR", user_id: $user_id, jwt: $token}')" --header 'Content-Type: application/json' http://localhost:8080/api/chirps)
# add_chirp_resp=$(curl -w "\n" -H "Authorization: Bearer $token" --data "$(jq -n --arg user_id "$user_id" '{body: "FOO BAR", user_id: $user_id}')" --header 'Content-Type: application/json' http://localhost:8080/api/chirps)
# add_chirp_resp=$(curl -w "\n" -H "Authorization: Bearer $token" --data "$(jq -n --arg user_id "$user_id" '{body: "FOO BAR", user_id: $user_id}')" --header 'Content-Type: application/json' http://localhost:8080/api/chirps)
# echo $add_chirp_resp
# formdata=$(jq -n --arg token "$token" --arg user_id "$user_id" '{body: "FOO BAR", user_id: $user_id, jwt: $token}')
# echo $formdata
