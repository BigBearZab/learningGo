package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// useArg()
	o, _ := os.Open(useArg())
	bs, _ := io.Copy(os.Stdout, o)
	fmt.Println("File byte size:", bs)
}

func useArg() string {
	return os.Args[1]
}
