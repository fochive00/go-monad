package main

import (
	"fmt"
	. "monad"
)

func main() {

	a := Just(100)

	b := AndThen(a, func(num int) Maybe[int] {
		return SafeDiv(num, 0)
	})

	c := AndThen(b, func(num int) Maybe[int] {
		return SafeDiv(num, 1000)
	})

	fmt.Println("c:", c)
}

func SafeDiv(a int, b int) Maybe[int] {
	if b == 0 {
		return Nothing[int]()
	}

	return Just(a / b)
}
