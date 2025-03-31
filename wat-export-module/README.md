# math

A simple module with exported, basic math functions in WebAssembly text format (`.wat`) file.

This folder demonstrates how to implement a simple module with basic math functions in WebAssembly text format (`.wat`) file.
It also demonstrates how to compile the `.wat` file to the binary format `.wasm` file, and run that either in the browser,
or in the console using Node.js.

Compile the Wasm text file to binary:

```bash
./build.sh
```

Run it, from Node.js:

```bash
node index.js

<Buffer 00 61 73 6d 01 00 00 00 01 06 01 60 01 7f 01 7f 03 02 01 00 07 0a 01 06 73 71 75 61 72 65 00 00 0a 09 01 07 00 20 00 20 00 6c 0b>
121
```

Run it in browser:

Start the HTTP server:

```bash
./server.sh
```

Open the  [http://localhost:8080](http://localhost:8080) URL link in a browser, then see the developer console (`Ctrl+Shift+J`).


Run with external runtime:

```bash

  $ wasmtime --invoke add math.wasm 42 24
  66

  $ wasmtime --invoke square math.wasm 2
  4

  $ wasmtime --invoke subtract math.wasm 42 24
  18

```
