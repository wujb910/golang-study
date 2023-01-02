package main

// func main() {
// 	println("Hello world")
// }

// package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("go version: %s", runtime.Version())
}
