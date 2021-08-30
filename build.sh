#!/bin/bash
echo "Building..."

cd ./src
GOOS=js GOARCH=wasm go build -v -o ../bin/hello.wasm
cp -u ./electron/* ../bin
cd ..

cp -u $(go env GOROOT)/misc/wasm/wasm_exec.js ./bin

cp -u ./assets/* ./bin
