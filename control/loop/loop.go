package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {

	}
	var i int
	fmt.Println(i)
	// for i < 10 {
	// }
	// for {
	// }
	// fmt.Println("aaa")

here:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if i == 0 {
				continue here // 退出当前循环
			}
			fmt.Println(i, j)
			if j == 2 {
				break
			}
		}
	}

there:
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if j == 1 {
				continue
			}
			fmt.Println(i, j)
			if j == 2 {
				break there // 退出整个循环
			}
		}
	}
}
