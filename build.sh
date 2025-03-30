#!/bin/bash
rm -rf ./build
mkdir ./build
GOOS=linux GOARCH=amd64 go build -o build/live_server_linux_amd64 main.go
GOOS=linux GOARCH=arm64 go build -o build/live_server_linux_arm64 main.go
GOOS=linux GOARCH=arm go build -o build/live_server_linux_arm main.go
GOOS=darwin GOARCH=arm64 go build -o build/live_server_mac_arm64 main.go
GOOS=windows GOARCH=amd64 go build -o build/live_server_win_amd64 main.go

chmod 755 ./build/live_server_linux_amd64
chmod 755 ./build/live_server_linux_arm64
chmod 755 ./build/live_server_linux_arm
chmod 755 ./build/live_server_mac_arm64
chmod 755 ./build/live_server_win_amd64
