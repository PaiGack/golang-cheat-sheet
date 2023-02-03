package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))
	fmt.Println(scope()())
	o, v := outer()
	fmt.Println(o(), v, o())

	o2, v2 := outer()
	fmt.Println(o2(), v2, o2())
}

func scope() func() int {
	outer_var := 2
	return func() int {
		return outer_var
	}
}

// func another_scope() func() int {
// 	outer_var = 444
// 	return foo
// }

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	return inner, outer_var
}
