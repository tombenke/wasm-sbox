#!/bin/bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
GOOS=js GOARCH=wasm tinygo build -o main_tinygo.wasm main.go

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./wasm_exec.js
cp /usr/local/lib/tinygo/targets/wasm_exec.js ./wasm_exec_tinygo.js

