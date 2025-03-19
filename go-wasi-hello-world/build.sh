#!/bin/bash

GOARCH=wasm GOOS=wasip1 go build -o main.wasm
GOARCH=wasm GOOS=wasip1 go build -buildmode=c-shared -o main_shared.wasm
