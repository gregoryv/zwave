#!/bin/bash

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
	gofmt -w $path
        go test .
	if [ $? == 0 ]; then
	    go build -o tailacm/tailacm zwave/tailacm
	    go build -o sendacm/sendacm zwave/sendacm
	    killall tailacm
	fi
        ;;
esac

