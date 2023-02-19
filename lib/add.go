package lib

import f "fmt"

func Add(a int, b int) int {
	return a + b
}

func init() {
	f.Println("init")
}
