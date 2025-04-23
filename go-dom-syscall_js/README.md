# go-dom-syscall_js

Build the Wasm files:

```bash
./build.sh
```

The results will be:

- `main_go.wasm`,
- `main_tinygo.wasm`.

Use the built modules in a Node.js application:

```bash
node index_nodejs_go.js 
add result:  66
sub result:  66
mul result:  1008

node index_nodejs_tinygo.js 
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
