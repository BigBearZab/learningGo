package main

import (
	"fmt"
	"log"
)

func main() {
	_, s := isPrime(7)
	log.Println(s)
}

func isPrime(n int) (bool, string) {
	if n <= 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}

	for i := 2; i <= n/2; i++ {
		if n%2 == 0 {
			return false, fmt.Sprintf("%d is not prime as it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is considered prime", n)
}
