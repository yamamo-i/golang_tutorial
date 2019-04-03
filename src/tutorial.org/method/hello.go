package test

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scala(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("before: %+v, Abvs: %v\n", v, v.Abs())
	v.Scala(5)
	fmt.Printf("after: %+v, Abvs: %v\n", v, v.Abs())
}
