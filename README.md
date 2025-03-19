# wasm-sbox

This is a WebAssembly sandbox.
It provides very simple examples about how to create WebAssembly modules and applications using different
programming languages, and how to use the created Wasm binaries within different environments and platforms.

## The list of examples

### Basics

- [c-emscripten-hello-world](c-emscripten-hello-world):
  Run a simple "Hello World" application in a HTML page that is written in C.

- [go-export-module](go-export-module):
  Exports some go functions, and build as a module, then call the functions in several environments:
  in a runner (e.g. [wasmtime](https://wasmtime.dev/)), in a Node.js application, in a HTML page.
    
- [go-export-module-syscall_js](go-export-module-syscall_js):
  Exports some go functions that uses the [`syscall/js`](https://pkg.go.dev/syscall/js) Go package, and build as a module,
  then call the functions in a Node.js application, and in a HTML page.

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

- [The ByteCode Alliance](https://bytecodealliance.org/)
- [WebAssembly](https://webassembly.org/)
- [WebAssembly Docs](https://developer.mozilla.org/en-US/docs/WebAssembly)
- [Web Assemply Core Specification](https://webassembly.github.io/spec/core/)
- [Awesome WebAssembly Languages](https://github.com/appcypher/awesome-wasm-langs)
- [Wasmer](https://wasmer.io/)
- [Wasmtime](https://wasmtime.dev/)
- [WABT: The WebAssembly Binary Toolkit](https://github.com/WebAssembly/wabt)
- [wapm - Web Assembly Package Manager](https://github.com/wasmerio/wapm-cli)
- [Web Platform]()
- [Web Assembly Component Model](https://component-model.bytecodealliance.org/)
- [Go Wiki: WebAssembly](https://go.dev/wiki/WebAssembly)
- [Extensible Wasm Applications with Go](https://go.dev/blog/wasmexport)
- [WebAssembly Directives](https://pkg.go.dev/cmd/compile#hdr-WebAssembly_Directives)
- [WASI support in Go](https://go.dev/blog/wasi)
- [The Go "syscall/js" package](https://pkg.go.dev/syscall/js)
- [TinyGo](https://tinygo.org/)
- [Using WASM - How to call WebAssembly from JavaScript in a browser](https://tinygo.org/docs/guides/webassembly/wasm/)
- [Emscriptenv](https://emscripten.org/)
  a complete compiler toolchain to WebAssembly, using LLVM, with a special focus on speed, size, and the Web platform.
- [Node.js with WebAssembly](https://nodejs.org/en/learn/getting-started/nodejs-with-webassembly).
