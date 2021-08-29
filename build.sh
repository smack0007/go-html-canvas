#!/bin/bash
echo "Building..."

cd ./src
go env -w GOOS=js GOARCH=wasm && go build -v -o ../bin/hello.wasm
cp ./electron/* ../bin
cd ..

cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./bin
