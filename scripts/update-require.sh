#!/bin/bash
# 移除依赖项版本锁定
rm go.sum
# 重新拉取依赖
go mod tidy