package main

import "fmt"

func functionName() {}

func functionName2(param1 string, param2 int) {}

func functionName3(param1, param2 int) int {
	return 1024
}

func functionName4() int {
	return 1024
}

func returnMulti() (int, string) {
	return 42, "foobar"
}

var x, str = returnMulti()

func returnMulti2() (n int, s string) {
	n = 42
	s = "foobar"
	return
}

var x2, str2 = returnMulti2()

func main() {
	functionName()
	functionName2("", 1)
	functionName3(1, 2)
	functionName4()

	fmt.Println(x, str)
	fmt.Println(x2, str2)

}
