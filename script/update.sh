#!/bin/bash

# 查找所有 .proto 文件，排除 idl/api.proto
find idl -type f -name "*.proto" ! -path "idl/api.proto" | while read -r proto_file; do
  # 对每个 .proto 文件执行 hz update 命令
  hz update -I idl -idl "$proto_file"
done

go mod tidy
