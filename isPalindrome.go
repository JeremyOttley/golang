package main

import (
	"fmt"
	"os"
)

func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}

func main() {

	var args = os.Args[1:]

	fmt.Println(isPalindrome(args[0]))

}
