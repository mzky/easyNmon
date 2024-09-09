#!/bin/bash
version=$(git log --date=iso --pretty=format:"%h@%cd" -1)
compile="$(date '+%Y-%m-%d_%H:%M:%S') By $(go version|awk '{print $3}')"
cat <<EOF | gofmt >common/version.go
package common

const (
    Version = "$version"
    Compile = "$compile"
)
EOF
go mod tidy
go build -ldflags "-w -s" -o easyNmon-x86_64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build -ldflags "-w -s" -o easyNmon-aarch64 main.go
upx easyNmon-x86_64 easyNmon-aarch64
