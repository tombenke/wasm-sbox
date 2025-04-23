package main

import (
	//	"syscall/js"
	"fmt"
	dom "honnef.co/go/js/dom/v2"
)

//func add(this js.Value, args []js.Value) interface{} {
//	return uint32(args[0].Int()) + uint32(args[1].Int())
//}
//
//func sub(this js.Value, args []js.Value) interface{} {
//	return uint32(args[0].Int()) - uint32(args[1].Int())
//}
//
//func mul(this js.Value, args []js.Value) interface{} {
//	return uint32(args[0].Int()) * uint32(args[1].Int())
//}

func makeP() {
	body := dom.GetWindow().Document().GetElementsByTagName("body")[0]
	println(fmt.Sprintf("Document.Body: %+v\n", body))
	body.SetAttribute("bgcolor", "#cdas6f")
}

func main() {
	makeP()
	//	js.Global().Set("add", js.FuncOf(add))
	//	js.Global().Set("sub", js.FuncOf(sub))
	//	js.Global().Set("mul", js.FuncOf(mul))

	<-make(chan bool)
}
