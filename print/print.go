package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,", "你好")
	p := struct{ X, Y int }{17, 2}
	fmt.Println("My point:", p, "x cooed=", p.X)
	s := fmt.Sprintln("My point:", p, "x cooed=", p.X)
	fmt.Println(s)

	fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e\n", 17, 17, 17, 17.0, 17.0)
	s2 := fmt.Sprintf("%d %f\n", 17, 17.0)
	fmt.Println(s2)

	hellomsg := `
	"hello" in 
	
	"我和你"`
	fmt.Println(hellomsg)
}
