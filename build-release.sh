#!/usr/bin/env bash
env GOOS=linux GOARCH=arm go build -o go-state-store_linux_armhf
env GOOS=linux GOARCH=amd64 go build -o go-state-store_linux_amd64
