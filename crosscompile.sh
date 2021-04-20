#!/bin/sh

# @see https://opensource.com/article/21/1/go-cross-compiling
# amd64 is compatible with x86_64
archs=(amd64 arm64 ppc64le ppc64 s390x)

for arch in ${archs[@]}; do
  env GOOS=linux GOARCH=${arch} go build -o bin/seraph_${arch}
done