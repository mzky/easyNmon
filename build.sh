#!/bin/bash
version=$(git log --date=iso --pretty=format:"%H @%cd" -1)
compile="$(date '+%Y-%m-%d %H:%M:%S') by $(go version)"
cat <<EOF | gofmt >common/version.go
package common

const (
    Version = "$version"
    Compile = "$compile"
)
EOF
go mod tidy
go build -ldflags "-w -s" -o easyNmon_amd64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GOARM=7 go build -ldflags "-w -s" -o easyNmon_arm64 main.go

