#!/bin/bash
go get github.com/mitchellh/gox
version=$(git log --date=iso --pretty=format:"%H @%cd" -1)
compile="$(date '+%Y-%m-%d %H:%M:%S') by $(go version)"
cat <<EOF | gofmt >common/version.go
package common

const (
    Version = "$version"
    Compile = "$compile"
)
EOF

gox -osarch="linux/amd64 linux/arm64" -ldflags "-w -s" ./...
mv main_linux_amd64 easyNmon_amd64
mv main_linux_arm64 easyNmon_arm64
upx easyNmon_amd64


