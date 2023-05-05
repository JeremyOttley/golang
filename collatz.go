package main

import (
	"fmt"
)

func gtZero(n int64) bool {
	return n >= 0
}

func isEven(n int64) bool {
	return n%2 == 0
}

func isOdd(n int64) bool {
	return !isEven(n)
}

func collatz(n int64) int64 {

	if isEven(n) {
		return n / 2
	}

	if isOdd(n) {
		return n*3 + 1
	} else {
		return n
	}
}

func main() {

	fmt.Println("Enter an integer: ")

	var n int64

	fmt.Scanln(&n)

	fmt.Printf("%d", collatz(n))
}
