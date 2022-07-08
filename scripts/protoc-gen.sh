#!/bin/bash
#./protoc --js_out=import_style=commonjs,binary:./javascript *.proto
#./protoc --go_out=. --go_opt=paths=source_relative ./internal/resources/user/*.proto
./protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    */*.proto \
    */*/*.proto \
    */*/*/*.proto \