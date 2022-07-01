#!/bin/bash

echo "Start build......"

# 构建项目
go build ./cmd/app/main.go

# 复制打包后文件
mv main ./deploy/main

# 复制配置文件
cp ./config/index.ini ./deploy/config/index.ini
cp ./config/pro.ini ./deploy/config/pro.ini

echo "End build."
