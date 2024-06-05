// Package main provides ...
package main

import (
	"bufio"
	"fmt"
	m "math" // Math library with local alias m.
	"os"
	"strconv"
)

// example of how to declare multiple constants at once.
const (
	// iota is used to create enumerated constants
	catSpecialist = iota // catSpecialist = 0
	dogSpecialist = iota // dogSpecialist = 1
	cimSpecialist = iota // cimSpecialist = 2
)

// example of how to declare multiple constants at once.
const (
	// iota is used to create enumerated constants
	zero = iota //  zero 	=  0
	one         //  one 	=  1
	two         //  two 	=  2
)

// example of how to declare multiple constants at once.
const (
	// _ is a wright only variable meaning you dont
	//care about it.
	_     = iota //  ignore the 0
	One          //  One =  1
	Two          //  Two =  2
	Three        //  Three =  3
)

/*
Rules: 1. use camelCase or Pascal dont use snake_case
	   if acronym use ALLCAPS.
	   2. when adding function comments make them a full
	   sentance and start with the function name
*/

func main() {
	// usingImports()
	// strings()
	// types()
	// typeConvert()
	// a := make([]int, 3)
	// passingArrays(a, 1)
	// mapExample()
	// arrayExample()
	// sliceFunc()
	// forLoop()
	// array := [...]int{1, 2, 3}
}

//StdioHandling
func StdioHandling() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println("line", s.Text())
	}
}

// INTERFACES

//learnInterfaces
func learnInterfaces() {
	// INTERFACES
	// Struct Declaration
	// name := StructType{value}
	// name := StructType{x: value}
	// name := StructType{x: value, y: value2}
	p := Num{x: 1}
	p.Print()
	p.Type()
	printDat(p)
}

//printDat
func printDat(p Printer) {
	p.Print()
	p.Type()
}

type Printer interface {
	Print()
	Type()
}

type Num struct {
	x int
}

func (i Num) Print() {
	fmt.Printf("i is %v\n", i.x)
}

// INTERFACES

//Type
func (i Num) Type() {
	fmt.Printf("i's type is %T\n", i.x)
}

// usingImports example of how using an imported library
func usingImports() {
	// using a function from math
	fmt.Println(m.Exp(10))
}

func strings() {
	str := "I am a string"
	str2 := " I am another string"
	fmt.Printf("%v, %T\n", str, str) // I am a string, string

	// Strings Act Like Arrays Of Bytes
	fmt.Printf("%v, %T\n", str[2], str[2])                 // 97, unint8
	fmt.Printf("%v, %T\n", string(str[2]), string(str[2])) // a, string

	// Concatenation
	// using + operator
	fmt.Printf("%v, %T\n", str+str2, str+str2) // I am a string I am another string, string
}

func types() {
	// All Variables Default To 0
	var a int // 0
	fmt.Println(a)
	var b string // ""
	fmt.Println(b)
	var c bool // false
	fmt.Println(c)

	// Different Sizes Working Together
	var x int = 10
	var y int8 = 10
	// go will not convert sizes automatically
	fmt.Println(x + int(y))

	// bytes
	var num byte = 1 << 2 // 100
	// printing in binary
	fmt.Printf("%b\n", num)
}

// typeConvert demos how to convert between types
func typeConvert() {
	// Note: type conversion is more like a function
	// complex printf
	fl := 42.0
	fmt.Printf("value fl is is %v, type is %T\n", fl, fl)

	// convert from float to int
	it := int(fl)
	fmt.Printf("value it is is %v, type is %T\n", it, it)

	// convert int to string
	// note that string conversion requires "strconv"
	st := strconv.Itoa(it)
	fmt.Printf("value st is is %v, type is %T\n", st, st)
}

//multiReturnTest demos how to accept multiple return statments
func multiReturnTest() {
	sum, prod := multiReturn(1, 2)
	fmt.Printf("sum is %v, prod is %v\n", sum, prod)
}

// multiReturn utilizes multiple returned values
// and multiple input veriables.
func multiReturn(x, y int) (sum, prod int) {
	// Alternatively you could simply return a type
	// eg func multiReturn(x, y int) (int, int) {

	// how to receives multiple returns
	// sum, prod := multiReturn(1, 2)
	return x + y, x * y // return 2 values
}

// passingArrays demos how to pass arrays into a function.
func passingArrays(nums []int, i int) (size int) {
	for ; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	return len(nums)
}

// mapExample shows the use of maps in go.
func mapExample() {
	mp := map[string]int{"three": 3, "four": 4}
	mp["one"] = 1
	fmt.Println(mp["one"])
}

// arrayExample demos the creation and use of arrays in go
func arrayExample() {
	// remember that these are a fixed length
	var a [4]int               // a = [ 0, 0, 0, 0 ]
	a2 := [...]int{1, 2, 3, 4} // a2 = [ 1, 2, 3, 4 ]
	b := a                     // b is a copy of a
	b2 := &a                   // b2 is a pointer to a
	a[1] = 9                   // set a to a = [ 0, 9, 0, 0 ] and *b2 to [ 0, 9, 0, 0 ]
	fmt.Printf(" b is %v, *b2 is %v\n", b, *b2)
	fmt.Printf(" a is %v, a2 is %v\n", a, a2)
}

// sliceFunc is a demo of how to use slices
func sliceFunc() {
	// A slice is simply a pointer to an array.
	// When the capacity is exceded a new array is created.
	// make(type, length[, capacity]) // capacity is optional
	a := make([]int, 2)  // [ 0, 0 ]
	a2 := []int{1, 2, 3} // [ 1, 2, 3 ]	// Note: [] for slice and  [...] makes a array
	var a3 []int         // this is the perfered way to create an empty slice
	fmt.Printf("a3 is %v\n", a3)
	fmt.Printf("a2 is %v\n", a2)
	b := a   // b and a now point to the same array
	a[0] = 1 // since a and b point to the same array they both are changed

	// APPENDING
	a = append(a, 3) // a now points to a new array = [ 1, 0, 3 ]
	// append can take as many values as you like
	a = append(a, 2, 9) // a with 2 and a 9 on the end = [ 1, 0, 3, 2, 9 ]
	// ... can be used to turn the slice []int{1,2,3} into 1, 2, 3
	a = append(a, []int{1, 2, 3}...) // a = [ 1, 0, 3, 2, 9, 1, 2, 3 ]

	a[1] = 4 // Since a nolonger points to the same array as b only a is changed
	fmt.Printf("a is %v, b is %v\n", a, b)

	// Length is all the elements in a slice.
	// Capacity is the maximum size of the current array
	// that the slice is pointing to.
	fmt.Printf("a is %v, a's length is %v, a's capacity is %v\n", a, len(a), cap(a))

	// SUBSLICES
	// a = [1, 4, 3, 2, 9, 1, 2, 3]
	b = a[:2]  // turn b into a subslice of a aka [1, 4]
	b = a[1:3] // turn b into a subslice of a from index 1 to the 3rd location aka [4, 3]
	// removing an element from the middle of a slice
	a = append(a[:2], a[3:]...) // removing the 3rd element requires the use of append
	fmt.Printf("a is %v\n", a)  // a = [1 4 2 9 1 2 3]
}

// forLoop is a Simple demo of the many for loops in go.
func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Print("hello world\n")
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	// Infinite loop
	for {
		// never ending
	}
}
