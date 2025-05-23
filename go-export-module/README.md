# go-export-module

Exports some go functions, and build as a module, then call the functions in several environments:
in a runner (e.g. [wasmtime](https://wasmtime.dev/)), in a Node.js application, in a HTML page.

The [export.go](export.go) exports some simple functions, that are exported,
and are reacheable to an external Wasm caller.

IMPORTANT NOTE:
To execute `.wasm` in a browser, we’ll also need a HTML page with some JavaScript code to connect everything together.
An additional glue code is also need to be added to the page, which is delivered with the Go compiler, and called `wasm_exec.js`.

The same major Go version of the compiler and `wasm_exec.js` support file must be used together.
That is, if a `.wasm` file is compiled using Go version 1.N, the corresponding `wasm_exec.js`file
must also be copied from Go version 1.N. __Other combinations are not supported!__

The `wasm_exec.js` glue code shipped with your standard Go distribution,
usually found at `$(go env GOROOT)/lib/wasm/wasm_exec.js`.

In case of TinyGo it is `/usr/local/lib/tinygo/targets/wasm_exec.js`.

The [build.sh](build.sh) script compiles it to `.wasm` format,
using both the standard Go compiler and the TinyGo compiler.
It also copies the required `wasm_exec.js` files to both the standard Go compiler as well as to the TinyGo one.

NOTE: The build uses the `-buildmode=c-shared` parameter.

Build the Wasm files:

```bash
./build.sh
```

The results will be:

- `export_wasi_go.wasm`,
- `export_wasi_tinygo.wasm`.

Call the exported function with wasmtime both with the go and the tinygo versions:

```bash
wasmtime run --invoke add export_wasi_go.wasm 42 24
66

wasmtime run --invoke add export_wasi_tinygo.wasm 42 24
66
```

Use the built modules in a Node.js application:

```bash
node index_nodejs_go.js 
<Buffer 00 61 73 6d 01 00 00 00 00 f2 80 80 80 00 0a 67 6f 3a 62 75 69 6c 64 69 64 ff 20 47 6f 20 62 75 69 6c 64 20 49 44 3a 20 22 33 31 6a 72 64 77 79 68 6e ... 1603388 more bytes>
add result:  66
sub result:  66
mul result:  1008

node index_nodejs_tinygo.js 
<Buffer 00 61 73 6d 01 00 00 00 01 32 09 60 01 7f 00 60 00 00 60 02 7f 7f 01 7f 60 04 7f 7f 7f 7f 01 7f 60 02 7f 7f 00 60 00 01 7f 60 01 7f 01 7f 60 03 7f 7f ... 113322 more bytes>
add result:  66
sub result:  66
mul result:  1008
```

Start the HTTP server:

```bash
./server.sh
```

then open them in the corresponding HTML pages
e.g. [http://localhost:8080/index_browser_go.html](http://localhost:8080/index_browser_go.html),
and [http://localhost:8080/index_browser_tinygo.html](http://localhost:8080/index_browser_tinygo.html).
