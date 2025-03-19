# go-export-module

Exports some go functions, and build as a module, then call the functions in a runner,
such as [wasmtime](https://wasmtime.dev/).

The [export.go](export.go) exports some simple functions, that are exported, and are reacheable to an external Wasm caller.

The [build.sh](build.sh) script compiles it to `.wasm` format, using both the standard Go compiler and the TinyGo compiler.

NOTE: The build uses the `-buildmode=c-shared` parameter.

The results will be:

- `export_go.wasm`,
- `export_tinygo.wasm`.

Build the Wasm files:

```bash
./build.sh
```

Call the exported function with wasmtime both with the go and the tinygo versions:

```bash
wasmtime run --invoke add export_go.wasm 42 24
66

wasmtime run --invoke add export_tinygo.wasm 42 24
66
```
