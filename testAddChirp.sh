#!/usr/bin/env bash
curl -w "\n" --data '{ "email": "saul@bettercall.com" }' --header 'Content-Type: application/json' http://localhost:8080/api/users
curl -w "\n" --data '{ "body": "If youre committed enough, you can make any story work.", "user_id": "db3a6753-a571-436b-8b71-0af92186791a" }' --header 'Content-Type: application/json' http://localhost:8080/api/chirps
