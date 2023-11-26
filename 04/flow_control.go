package main

import (
	"fmt"
	"unicode/utf8"
)

func abc() {

	q := 1
	defer func() {
		q++
		fmt.Println("1", q)
	}()

	defer func() {
		q++
		fmt.Println("2", q)
	}()

	q = 100
	q++
	fmt.Println("3", q)
}

func Double(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println(v.(string) + v.(string))
	case int:
		fmt.Println(v.(int) * 2)
	default:
		fmt.Println(fmt.Sprintf("unsupport type=%T", v))
	}
}

func main() {
	abc()

	a := 0
	b := interface{}(a)

	fmt.Println(fmt.Sprintf("%T", b))

	Double("123")
	Double(123)
	Double(true)

	s := "abc发"
	for i, v := range s {
		fmt.Println(i, v, string(v))
	}

	fmt.Println(utf8.RuneCountInString(s), len(s), s[3:6])

}
