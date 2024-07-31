#!/bin/bash
version="v1.2.5"
project="misgo"

go mod tidy
go build -ldflags -s -ldflags -w -o $project .

docker build -t $project:$version .
rm -rf $project
echo "build success $project:$version"
