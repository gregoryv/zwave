#!/bin/bash

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
#	gofmt -w $path
        ;;
esac

go install github.com/gregoryv/zwave/cmd/tailacm
go install github.com/gregoryv/zwave/cmd/sendacm
go test .
