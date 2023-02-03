package main

import (
	"fmt"
	"time"
)

func doStuff(s string) {
	fmt.Println(s)
}

func main() {
	go doStuff("foobar")

	go func(x int) {
		fmt.Println(x)
	}(42)

	time.Sleep(time.Second * 10)
}
