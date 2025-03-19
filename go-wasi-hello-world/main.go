package main

import (
	"fmt"
)

//go:wasmexport AnExportedFunction
func AnExportedFunction() {
	println("This is an exported function call...")
}

//go:wasmexport AnExportedFunctionWithPar
func AnExportedFunctionWithPar(intValue int32) {
	fmt.Printf("This is an exported function call with parameter: %d\n", intValue)
}

func main() {
	println("Hello, WebAssembly")
}
