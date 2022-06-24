package main

import "fmt"

func main() {
	sl := GenerateSlice(10)

	for _, v := range sl {
		EvenOrOdd(v)
	}
}

func GenerateSlice(n int) []int {
	var sl []int
	for i := 0; i <= n; i++ {
		sl = append(sl, i)
	}
	return sl
}

func EvenOrOdd(n int) {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
