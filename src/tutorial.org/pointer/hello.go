package main

import (
	"fmt"
)

type Vertex struct {
	// X, Y int
	Lat, Long float64
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m = map[string]Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
