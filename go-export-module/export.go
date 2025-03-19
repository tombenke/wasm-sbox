package main

import ()

//go:wasmexport add
func add(x, y uint32) uint32 {
	return x + y
}

//go:wasmexport sub
func sub(x, y uint32) uint32 {
	return x + y
}

//go:wasmexport mul
func mul(x, y uint32) uint32 {
	return x * y
}

func main() {
	<-make(chan bool)
}
