package main

import (
	"fmt"
	"runtime"
)

func _if() int {
	x := 100
	if x > 10 {
		return x
	} else if x == 10 {
		return 10
	} else {
		return -1
	}
}

func _switch(operatingSystem string) {
	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
	case "linux":
		fmt.Println("Liunx Geek")
	default:
		fmt.Println("Other")
	}
}

func main() {
	_if()
	switch os := runtime.GOOS; os {
	case "":
		fmt.Printf("...")
	}
	_switch(runtime.GOOS)

	number := 42
	switch {
	case number < 42:
		fmt.Println("Smaller")
	case number == 42:
		fmt.Println("Equal")
	case number > 42:
		fmt.Println("Greater")
	}

	var char byte = '?'
	switch char {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Should escape")
	}
}
