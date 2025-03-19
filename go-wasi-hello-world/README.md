# go-wasi-hello-world

A go program that prints a message to the console within WASI,
using [wasmer](https://wasmer.io/) and [wasmtime](https://wasmtime.dev/) runners.

This application can be used both as an application (prints "Hello World") or a module, that exports some functions.

Build the binaries with with `_start` to use as an application, and without `_start`, to use as a module:

```bash
./build.sh
```

See the [bash.sh](bash.sh) for details about the command parameters.

Run the applitation with wasmer:

```bash
wasmer run main.wasm 
Hello, WebAssembly
```

or run with wasmtime:

```bash
wasmer run main.wasm 
Hello, WebAssembly
```


Run with calling an exported function, using wasmer and wasmtime:

```bash
wasmer run --invoke AnExportedFunction main_shared.wasm 
This is an exported function call...
```

Run with calling an exported function with parameter, using wasmtime:

```bash
wasmtime run --invoke AnExportedFunctionWithPar main_shared.wasm 42
This is an exported function call...
```

