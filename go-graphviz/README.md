# go-graphviz

## About

Draw a graph using Golang and Graphviz in an HTML page.

This application demonstrates how to create a graph is SVG format and visualize it in an HTML page.

The page loads two wasm modules:
    - graphviz [@hpcc-js/wasm](https://github.com/hpcc-systems/hpcc-js-wasm/),
    - the main application, that is written in Go and compiled into wasm binary.

The graphviz graph renderer is exported as a global function, using JavaScript.

The Go application uses the [GoGraphviz](https://github.com/mzohreva/GoGraphviz) package
in order to generate a graph in the form of the dot language of the Graphviz tool.

Finally the Go application calls the graphviz renderer with the dot content to convert it into SVG format,
then place it into the page.

## Usage

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

## References

- [GoGraphviz](https://github.com/mzohreva/GoGraphviz)
- [@hpcc-js/wasm](https://github.com/hpcc-systems/hpcc-js-wasm/)
