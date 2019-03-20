package main

import (
	"fmt"
	"os"

	"./file"
)

func main() {
	hello := []byte("hello world\n")
	file.Stdout.Write(hello)
	f, err := file.Open("/does/not/exitst")
	if f == nil {
		fmt.Printf("cannot opejn file;; err=%s\n", err.String())
		os.Exit(1)
	}
}
