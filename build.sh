#!/bin/bash
version="v1.3.82"
project="misgo"

go mod tidy
go build -ldflags -s -ldflags -w -o $project .

docker build -t $project:$version .
rm -rf $project
echo "build success $project:$version"
