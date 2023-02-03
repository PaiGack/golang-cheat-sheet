package main

import "fmt"

func main() {
	// var foo int
	// var foo int = 42
	// var foo, bar int = 42, 1302
	// var foo = 42
	// foo := 42
	const constant = "This is a constant"
	fmt.Println(constant)

	const (
		_ = iota
		a
		b
		_
		c = 1 << iota
		d
	)

	fmt.Println(a, b, c, d)

}
