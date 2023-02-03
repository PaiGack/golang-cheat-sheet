package main

import "fmt"

var cp64 complex64

func main() {
	cp64 = 12 + 1i
	cp64 = 1i
	cp64 = 12
	_ = real(cp64)
	fmt.Printf("cp64: %v\n", cp64*cp64)
	// cp64.

	var i int = -1
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(u)
}
