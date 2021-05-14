package main

import (
	"fmt"
	"io"
	"os"
)

type logger struct{}

func main () {
	fn := os.Args[1]
	f,err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logger{}
	io.Copy(lw, f)
}

func (logger) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
