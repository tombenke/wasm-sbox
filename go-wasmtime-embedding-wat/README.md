# go-wasmtime-embedding-wat

Run a WebAssembly text (`.wat`)format application embedded into a Go application from external source file,
using [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go).

This [example](main.go) demonstrates how to use the [wasmtime-go](https://github.com/bytecodealliance/wasmtime-go) runtime
built into a go application, that loads text format WebAssemly code from the [gcd.wat](gcd.wat) external file, then runs it.

Run the application:

```bash
go run main.go

gcd(6, 27) = 3
```
