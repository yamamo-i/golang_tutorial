package main

import (
	crnad "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

func main() {
	seed, _ := crnad.Int(crnad.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	fmt.Println("number is ", rand.Intn(100000))
}
