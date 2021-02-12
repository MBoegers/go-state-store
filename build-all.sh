#!/usr/bin/env bash
echo "Create directories"
mkdir -p target/bin

echo "Copy configuration"
cp configuration/default-config.yml target/default-config.yml

echo "Build binaries"
env GOOS=linux GOARCH=arm go build -o target/bin/go-state-store_linux_arm
env GOOS=linux GOARCH=arm64 go build -o target/bin/go-state-store_linux_arm64
env GOOS=linux GOARCH=amd64 go build -o target/bin/go-state-store_linux_amd64
env GOOS=linux GOARCH=386 go build -o target/bin/go-state-store_linux_386