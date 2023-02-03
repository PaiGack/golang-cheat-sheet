package main

import (
	"fmt"
	"time"
)

func main() {
	// var a []int
	// a = append(a, 1, 2, 3, 4)

	// var a = []int{1, 2, 3, 4}
	// a = append(a, a...)

	a := []int{1, 2, 3, 4}
	chars := []string{0: "a", 2: "c", 3: "e"}
	b := []int{1: 1, 2: 2, 10: 10}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(chars)

	// var c = a[1:5] run panic
	var c = a[1:4]
	c = append(c, 100)
	fmt.Println(c)
	fmt.Println(a)

	var c2 = a[:3]
	var c3 = a[3:]
	fmt.Println(c2, c3)

	d1 := make([]byte, 5, 5)
	d1 = make([]byte, 5)
	d1 = make([]byte, 0, 5)
	d1 = append(d1, 1)
	fmt.Println(d1)

	x := [3]string{"abc", "def", "hij"}
	s := x[:]
	fmt.Println(x, s)
	x[2] = "h"
	s[0] = "a"
	s = append(s, "mni")
	s[1] = "d"
	fmt.Println(x, s)

	for i, e := range a {
		fmt.Println(i, e)
	}

	for _, e := range a {
		fmt.Println(e)
	}

	for i := range a {
		fmt.Println(i)
	}

	// 循环 a 的长度
	for range a {
		fmt.Println(a)
	}

	add(a)
	fmt.Println(a)

	// 每秒执行一次
	for range time.Tick(time.Second) {
		fmt.Println("1s")
	}
}

func add(a []int) {
	a[0] = 10086
	fmt.Println(a)
}
