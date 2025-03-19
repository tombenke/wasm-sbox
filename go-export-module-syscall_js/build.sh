#!/bin/bash
GOOS=js GOARCH=wasm go build -o export_js_go.wasm export_syscall_js.go
GOOS=js GOARCH=wasm tinygo build -o export_js_tinygo.wasm export_syscall_js.go

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./wasm_exec_go.js
cp /usr/local/lib/tinygo/targets/wasm_exec.js ./wasm_exec_tinygo.js

