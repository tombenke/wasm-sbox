# go-browser-hello-world

A WebAssembly application written in Go, that applies the [`syscall/js`](https://pkg.go.dev/syscall/js) package to use the DOM API.

Build the application:

```bash
./build.sh
```

Start the HTTP server:

```bash
./server.sh
```

Then open either the [http://localhost:8080/index_go.html](http://localhost:8080/index_go.html) URL,
or the [http://localhost:8080/index_tinygo.html](http://localhost:8080/index_tinygo.html).
