package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	return uint32(args[0].Int()) + uint32(args[1].Int())
}

func sub(this js.Value, args []js.Value) interface{} {
	return uint32(args[0].Int()) - uint32(args[1].Int())
}

func mul(this js.Value, args []js.Value) interface{} {
	return uint32(args[0].Int()) * uint32(args[1].Int())
}

func main() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("sub", js.FuncOf(sub))
	js.Global().Set("mul", js.FuncOf(mul))
	<-make(chan bool)
}
