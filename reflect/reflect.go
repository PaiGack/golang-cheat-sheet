package main

import (
	"fmt"
)

func main() {
	do(21)
	do("Hello")
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v %v\n", v, v*2, i)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don,t konw about type %T!\n", v)
	}
}
