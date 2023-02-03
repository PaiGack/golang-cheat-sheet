package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() { ch <- 42 }()
	v := <-ch
	fmt.Println(v)

	ch2 := make(chan int, 100)
	close(ch2)
	v, ok := <-ch2
	if !ok {
		fmt.Println(v, ok)
	}

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

	chIn, chOut := make(chan int), make(chan int)
	// var cc chan int nil chan
	// go func() { cc <- 1 }()
	// x := <-cc
	// fmt.Println(x)

	go func() { _ = <-chOut }()
	go func() { chIn <- 10 }()
	doStuff(chOut, chIn)

	// // fail
	// var c chan string
	// c <- "Hello world!"

	// // fial
	// var c chan string
	// fmt.Println(<-c)

	// // panic
	// var c = make(chan string, 1)
	// c <- "Hello, World!"
	// close(c)
	// c <- "Hello, World!"

	var c = make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	for i := 0; i < 3; i++ {
		fmt.Println(<-c) // 依次读出 1 2 0
	}
}

func doStuff(channelOut, channelIn chan int) {
	select {
	case channelOut <- 42:
		fmt.Println("We could write to channelOut!")
	case x := <-channelIn:
		fmt.Printf("We chould read from channelIn %v\n", x)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
