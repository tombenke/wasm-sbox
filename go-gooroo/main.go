package main

import (
	"fmt"
	o "github.com/Matbabs/Gooroo"
	dom "github.com/Matbabs/Gooroo/dom"
)

// Link struct holds the properties of a Web URL
type Link struct {
	Title string
	URL   string
}

// The list of Links to display
var Links = []Link{
	Link{Title: "The ByteCode Alliance", URL: "https://bytecodealliance.org/"},
	Link{Title: "WebAssembly", URL: "https://webassembly.org/"},
	Link{Title: "WebAssembly Docs", URL: "https://developer.mozilla.org/en-US/docs/WebAssembly"},
	Link{Title: "Web Assemply Core Specification", URL: "https://webassembly.github.io/spec/core/"},
	Link{Title: "Awesome WebAssembly Languages", URL: "https://github.com/appcypher/awesome-wasm-langs"},
	Link{Title: "Wasmer", URL: "https://wasmer.io/"},
	Link{Title: "Wasmtime", URL: "https://wasmtime.dev/"},
	Link{Title: "WABT: The WebAssembly Binary Toolkit", URL: "https://github.com/WebAssembly/wabt"},
	Link{Title: "wapm - Web Assembly Package Manager", URL: "https://github.com/wasmerio/wapm-cli"},
	Link{Title: "Web Platform", URL: "https://webplatform.github.io/"},
	Link{Title: "Web APIs", URL: "https://developer.mozilla.org/en-US/docs/Web/API"},
	Link{Title: "Web Assembly Component Model", URL: "https://component-model.bytecodealliance.org/"},
	Link{Title: "Go Wiki: WebAssembly", URL: "https://go.dev/wiki/WebAssembly"},
	Link{Title: "Extensible Wasm Applications with Go", URL: "https://go.dev/blog/wasmexport"},
	Link{Title: "WebAssembly Directives for Go", URL: "https://pkg.go.dev/cmd/compile#hdr-WebAssembly_Directives"},
	Link{Title: "WASI support in Go", URL: "https://go.dev/blog/wasi"},
	Link{Title: "The Go \"syscall/js\" package", URL: "https://pkg.go.dev/syscall/js"},
	Link{Title: "TinyGo", URL: "https://tinygo.org/"},
	Link{Title: "Using WASM - How to call WebAssembly from JavaScript in a browser", URL: "https://tinygo.org/docs/guides/webassembly/wasm/"},
	Link{Title: "Emscriptenv - complete compiler toolchain to WebAssembly, using LLVM, with focus on speed, size, and the Web platform.", URL: "https://emscripten.org/"},
	Link{Title: "Node.js with WebAssembly", URL: "https://nodejs.org/en/learn/getting-started/nodejs-with-webassembly."},
}

// Declare an attribute of an html element with the value 'target='
func HrefTarget(target string) o.DomComponent {
	//TODO: Add sanitizeHtml(&target)
	return func() string {
		return fmt.Sprintf("%s%s'%s'", dom.ELEMENT_PARAM, "target=" /*TODO: Add dom.HTML_PARAM_TARGET*/, target)
	}
}

// WaslLinks creates a DomComponent that holds the list of links
func WasmLinks() o.DomComponent {
	return o.Div(
		o.P("WebAssembly links:"),
		o.Ul(
			o.For(Links, func(i int) o.DomComponent {
				return o.Li(o.A(Links[i].Title, HrefTarget("_empty"), o.Href(Links[i].URL)))
			}),
		),
	)
}

// App defines the main component of the application
func App() o.DomComponent {

	return o.Div(
		o.H1("Hello WebAssembly and Gooroo!"),
		o.Div(
			WasmLinks(),
		),
	)

}

func main() {
	o.Render(func() {
		o.Html(
			App(),
		)
	})
}
