#!/bin/bash
GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o export_wasi_go.wasm export.go
GOOS=wasip1 GOARCH=wasm tinygo build -buildmode=c-shared -o export_wasi_tinygo.wasm export.go

GOOS=js GOARCH=wasm go build -o export_js_go.wasm export.go
GOOS=js GOARCH=wasm tinygo build -o export_js_tinygo.wasm export.go

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" ./wasm_exec_go.js
cp /usr/local/lib/tinygo/targets/wasm_exec.js ./wasm_exec_tinygo.js
