package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		y := z - (z*z-x)/(2*z)
		if math.Abs(z-y) < 1.e-8 {
			break
		}
		z = z - y
	}
	return math.Abs(z)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("far away.....orz")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morninig!")
	case t.Hour() < 17:
		fmt.Println("Goot afternoon")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("delay world")
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("counting")
}
