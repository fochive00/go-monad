package main

import (
	"fmt"
	. "monad"
)

func main() {

	a := Just(2)
	b := Monad[int, int](a)(CurrySafeDiv(4))
	c := Monad[int, int](b)(CurrySafeDiv(8))
	d := Monad[int, int](c)(CurrySafeDiv(3))
	e := Monad[int, int](d)(CurrySafeDiv(16))

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	f := Monad[int, int](Monad[int, int](Monad[int, int](Monad[int, int](a)(
		CurrySafeDiv(4)))(
		CurrySafeDiv(8)))(
		CurrySafeDiv(3)))(
		CurrySafeDiv(16))

	fmt.Println("f:", f)
}

func SafeDiv(a int, b int) Maybe[int] {
	if b == 0 {
		return Nothing[int]()
	}

	return Just(a / b)
}

func CurrySafeDiv(a int) func(b int) Maybe[int] {
	return func(b int) Maybe[int] {
		if b == 0 {
			return Nothing[int]()
		}

		return Just(a / b)
	}
}
