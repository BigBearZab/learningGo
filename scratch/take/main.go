package main

import (
	"fmt"
)

func main() {
	a1 := []string{"apple", "pear", "lemon", "orange"}
	n1 := 8
	fmt.Println(takeMed1(a1, n1))
	fmt.Println(n1 / len(a1))
}

func takeEasy(a []string, n int) []string {
	if n > len(a) {
		return a
	}

	return a[:n]
}

func takeMed1(a []string, n int) []string {
	whole := n / len(a)
	rem := n % len(a)
	var out_slice []string
	for i := 1; i <= whole; i++ {

		out_slice = append(out_slice, a...)
	}
	out_slice = append(out_slice, a[:rem]...)
	return out_slice
	// append(out_slice, a)
}
