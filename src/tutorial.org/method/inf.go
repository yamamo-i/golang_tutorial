package inf

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// var i I

	// var t *T
	// i = t
	// describe(i)
	// i.M()

	// i = &T{"hello"}
	// describe(i)
	// i.M()

	// i = F(math.Pi)
	// describe(i)
	// i.M()

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
