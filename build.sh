#!/bin/bash
echo "Building..."
go env -w GOOS=js GOARCH=wasm && go build -o ./bin/hello.wasm ./src
cp ./src/electron/* ./bin
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./bin
