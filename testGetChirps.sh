#!/usr/bin/env bash
# curl --request POST http://localhost:8080/admin/reset
# curl --data '{"name":"bob"}' --header 'Content-Type: application/json' http://example.com/users/1234
# curl --data '{ "email": "saul@bettercall.com" }' --header 'Content-Type: application/json' http://localhost:8080/api/users
curl -w "\n" --data '{ "body": "If youre committed enough, you can make any story work.", "user_id": "bfc7ea05-9c51-4664-ad54-5218c8fa94ed" }' --header 'Content-Type: application/json' http://localhost:8080/api/chirps
# curl http://localhost:8080/api/chirps
