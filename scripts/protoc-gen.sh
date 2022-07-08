#!/bin/bash

# 根据.proto生成.pb.go文件
./protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    */*.proto \
    */*/*.proto \
    */*/*/*.proto \

# 根据.proto生成.pb.js文件
./protoc --js_out=import_style=commonjs,binary:./ ./*/*.proto
./protoc --js_out=import_style=commonjs,binary:./ ./*/*/*.proto
./protoc --js_out=import_style=commonjs,binary:./ ./*/*/*/*.proto

# 为生成的.pb.go文件注入自定义tag
./protoc-go-inject-tag \
    -input=*/*.pb.go \
    -input=*/*/*.pb.go \
    -input=*/*/*/*.pb.go \