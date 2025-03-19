# go-wasmtime-embedding-wat-inline

Run a WebAssembly text (`.wat`)format application embedded into a Go application,
using [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go).

The [main.go](main.go) uses the [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go) package
to run a text format Wasm source file in a [wasmtime](https://github.com/bytecodealliance/wasmtime) runtime,
that is embedded into a Go application.

Run the example:

```bash
go run main.go

Hello from Go!
```
