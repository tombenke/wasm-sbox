package main

import (
	"syscall/js"
)

func main() {
	println("Hello, WebAssembly")

	message := "👋 Hello World 🌍"

	document := js.Global().Get("document")
	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", message)
	document.Get("body").Call("appendChild", h2)

	<-make(chan bool)
}
