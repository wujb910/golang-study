package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(8, 4)
	fmt.Println(result, err)
	result, err = divide(8, 0)
	fmt.Println(result, err)
	r, e, op := operate(8, 0, divide)
	fmt.Println(r, e, op, divide)
	fmt.Println(op(9, 3))

}

func multiplication(a, b int) int {
	return a * b
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("除数不能为0")
	}

	return x / y, nil
}

type divideOperation func(m, n int) (int, error)

func operate(o, p int, op divideOperation) (int, error, divideOperation) {
	if op == nil {
		return 0, errors.New("op is nil"), nil
	}
	r, e := op(o, p)
	return r, e, op
}
