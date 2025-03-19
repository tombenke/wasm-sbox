#!/bin/bash

docker run --rm -v $(pwd):/src -u $(id -u):$(id -g) emscripten/emsdk emcc hello.c -o hello.html

