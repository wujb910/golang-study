package main

import "fmt"

func main() {
	const (
		red = 1 << iota
		yellow
		blue
	)

	type colourPaper struct {
		colours int8
	}

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
