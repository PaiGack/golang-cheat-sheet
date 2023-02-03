package main

import (
	"fmt"
	"log"
	"math"
)

type Vertex struct {
	X, Y float64
}

var v = Vertex{1, 2}
var v2 = Vertex{X: 1, Y: 2}
var v3 = []Vertex{{1, 2}, {3, 4}}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) add(n float64) {
	v.X += n
	v.Y += n
}

type Awesomizer interface {
	Awesome() string
}

type Foo struct{}

func (foo Foo) Awesome() string {
	return "Awesome"
}

func main() {
	v.X = 100
	fmt.Println(v, v.X)
	fmt.Println(v.Abs())
	v.add(100)
	fmt.Println(v)

	point := struct{ X, Y int }{1, 2}
	fmt.Println(point)

	p := Vertex{1, 2}
	q := &p
	r := &Vertex{1, 2}
	q.X = 100
	var s *Vertex = new(Vertex)
	s.X = 100
	fmt.Println(q, r, s)

	foo := Foo{}
	fmt.Println(foo.Awesome())

	type Reader interface{}
	type Writer interface{}

	type ReadWriter interface {
		Reader
		Writer
	}

	type Server struct {
		Host string
		Port int
		*log.Logger
	}

	server := &Server{"localhost", 80, log.Default()}
	server.Logger.Fatal("exit")

	var logger *log.Logger = server.Logger
	fmt.Println(logger)

}
