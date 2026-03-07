#!/usr/bin/env bash
curl -w "\n" --data '{ "email": "mloneusk@example.co" }' --header 'Content-Type: application/json' http://localhost:8080/api/users
