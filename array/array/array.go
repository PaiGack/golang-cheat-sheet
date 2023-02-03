package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a)
	a[3] = 42
	i := a[3]
	fmt.Println(i)

	// var b = [3]int{1, 2, 3, 3} go build error
	// var b = [3]int{1, 2}
	// b := [3]int{1, 2}
	b := [...]int{1, 2}

	// b = append(b, 1) go build error
	fmt.Println(b)

	add(a)
	fmt.Println(a)
}

func add(a [10]int) {
	a[1] = 100
	fmt.Println(a)
}
