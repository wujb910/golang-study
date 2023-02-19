package main

// func main() {
// 	println("Hello world")
// }

// package main

import (
	"fmt"
	"golang-study/lib"
	"runtime"
	"strconv"
)

func main() {
	s := "aaaa"
	// var s1 string = "100"

	n, err := strconv.Atoi(s)
	fmt.Println(n, err)

	fmt.Printf("go version: %s\n", runtime.Version())
	fmt.Println(lib.Add(1, 2))
}
