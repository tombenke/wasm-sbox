# emscripten-hello-world

Run a simple "Hello World" application in a HTML page that is written in C.

The [hello.c](hello.c) file implements the application.

The [build.sh](build.sh) script compiles it to `.wasm` format with [emscripten](https://emscripten.org/),
using the `emscripten/emsdk` docker image.

Build the Wasm files, and generate the supporting files:

```bash
./build.sh
```

This will generate the following outputs:

- [hello.wasm](hello.wasm): The binary Web Assembly code.
- [hello.html](hello.html): The HTML file that loads the Web Assembly module.
- [hello.js](hello.js): The JavaScript loader to the HTML page.

Run a static HTTP server:

```bash
./server.sh
025/03/17 17:48:17 Server listens on 8080 port...
```

Open the [http://localhost:8080/hello.html](http://localhost:8080/hello.html) URL to see the page,
and open the developer console (Ctrl+Shift+J) to see the message printed out by the code.

 
