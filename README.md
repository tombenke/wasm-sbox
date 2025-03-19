# wasm-sbox

This is a WebAssembly sandbox.
It provides very simple examples about how to create WebAssembly modules and applications using different
programming languages, and how to use the created Wasm binaries within different environments and platforms.

## The list of examples

### Basics

- [c-emscripten-hello-world](c-emscripten-hello-world):
  Run a simple "Hello World" application in a HTML page that is written in C.

- [go-export-module](go-export-module):
  Exports some go functions, and build as a module, then call the functions in a runner,
  such as [wasmtime](https://wasmtime.dev/).

- [go-wasi-hello-world](go-wasi-hello-world):
  A go program that prints a message to the console within WASI,
  using [wasmer](https://wasmer.io/) and [wasmtime](https://wasmtime.dev/) runners.

- [go-wasmtime-embedding-wat-inline](go-wasmtime-embedding-wat-inline):
  Run a WebAssembly text (`.wat`)format application embedded into a Go application as inline text,
  using [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go).

- [go-wasmtime-embedding-wat](go-wasmtime-embedding-wat):
  Run a WebAssembly text (`.wat`)format application embedded into a Go application from external source file,
  using [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go).

- [wat-empty-module](wat-empty-module):
  This is the bare-minimum WebAssembly text format (`.wat`) file, that contains a single, empty module definition.

- [wat-export-module](wat-export-module):
  A simple module with exported, basic math functions in WebAssembly text format (`.wat`) file.

- `server.go`:
  This is a very simple HTTP file server that serves the examples that need a server to make the HTML pages working.

## References


