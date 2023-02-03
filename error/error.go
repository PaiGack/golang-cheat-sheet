package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	return math.Sqrt(x), nil
}

func main() {
	val, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
