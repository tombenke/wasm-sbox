#!/bin/bash
GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o export_go.wasm export.go
GOOS=wasip1 GOARCH=wasm tinygo build -buildmode=c-shared -o export_tinygo.wasm export.go

