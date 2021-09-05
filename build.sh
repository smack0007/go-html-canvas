#!/bin/bash
echo "Building..."

(cd ./src/main && GOOS=js GOARCH=wasm go build -v -o ../../bin/hello.wasm)

cp -u ./src/electron/* ./bin

cp -u $(go env GOROOT)/misc/wasm/wasm_exec.js ./bin

cp -u ./assets/* ./bin
