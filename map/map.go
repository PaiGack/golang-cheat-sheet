package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)

	// var m2 map[string]int
	// fmt.Println(m2)
	// m2["key"] = 1 // 未初始化，不能使用

	delete(m, "key")
	fmt.Println(m)

	elem, ok := m["key"]
	fmt.Println(elem, ok)

	var m3 = map[string]int{
		"key": 42,
	}
	fmt.Println(m3)

	for key, value := range m3 {
		fmt.Println(key, value)
	}

	for key := range m3 {
		fmt.Println(key)
	}
}
