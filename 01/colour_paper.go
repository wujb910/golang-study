package main

import (
	"fmt"
	"golang-study/lib"
)

func main() {
	lib.Add(1, 2)
	const (
		red = 1 << iota
		yellow
		blue
	)

	zhiyin := make(map[string][]int)
	zhiyin["a"] = []int{12345}
	zhiyin["b"] = make([]int, 0)
	zhiyin["b"] = append(zhiyin["b"], 123456)

	// var a string = "\"a\""
	// var b string = `"a"`

	type colourPaper struct {
		colours int8
	}
	// x := colourPaper{
	// 	colours: red,
	// }
	// var paper [5]colourPaper
	paper := [5]colourPaper{
		{
			colours: 4,
		}, {
			colours: 6,
		}, {
			colours: 7,
		}, {
			colours: 5,
		}, {
			colours: 3,
		},
	}

	for i, v := range paper {
		if v.colours&yellow > 0 {
			fmt.Println(i)
		}
	}
}
