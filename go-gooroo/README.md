# go-gooroo

A [Gooroo](https://github.com/Matbabs/Gooroo) application.

[Gooroo](https://github.com/Matbabs/Gooroo) is a web frontend library in Go (WebAssembly).

The Gooroo package gathers a set of cumulative functions allowing you to create web applications on the Frontend side.
To do this purpose, it implements DOM manipulation features based on `syscall/js` and webassembly.
Its objective is to explore the possibilities of a modern, lightweight and javascript independent web library.

Build the Wasm files:

```bash
./build.sh
```

Start the HTTP server:

```bash
./server.sh
```

then open them in the corresponding HTML pages
e.g. [http://localhost:8080/index.html](http://localhost:8080/index_go.html),
and [http://localhost:8080/index_tinygo.html](http://localhost:8080/index_tinygo.html).
