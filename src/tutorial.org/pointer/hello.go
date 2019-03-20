package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	// i := 42
	// p := &i
	// fmt.Println(*p)
	// *p = 21
	// fmt.Println(i, *p)

	// var v Vertex = Vertex{1, 2}
	// var p *Vertex = &v
	// p.X = 1e9
	// fmt.Println(v)
	// p.X = 2
	// fmt.Println(v1, p, v2, v3)

	// var a [2]string
	// a[0] = "hello"
	// a[1] = "world"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// var s []int = primes[1:4]
	// fmt.Println(s)

	// names := [4]string{
	// 	"Jhon",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)
	// b[0] = "xxx"
	// fmt.Println(a, b)
	// fmt.Println(names)

	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
