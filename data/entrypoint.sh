#!/bin/sh

go run main.go &
ngrok http 1234 
#curl http://127.0.0.1:4040/api/tunnels
