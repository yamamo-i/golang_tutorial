package main

import "fmt"

func add(hoge, fuga int) int {
	return hoge + fuga
}

func hoge(x, y int) (int, int) {
	a := x + y
	m := x - y
	return a, m
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	add, minus := hoge(20, 30)
	fmt.Println(add, minus)
	fmt.Println(split(17))
	var c, python, java = true, false, "no!!"
	fmt.Println(i, c, python, java)
}
