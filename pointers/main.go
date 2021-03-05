package main

import "fmt"

/*
	Go is a pass by value language, means
	every time you pass a params or reciever
	to go func it will create a copy of it.

	So to clear this picture we got POINTER's in
	go.

	A pointer is a variable whose value is the
	address of another variable.

	& => memory location of the variable
	* => value at the memory location

	var ip *int        // pointer to an integer
	var fp *float32    // pointer to a float

*/

func main() {
	var a int64
	a = 10
	fmt.Printf("%d\n", a)
	addTenToANumber(a)
	fmt.Printf("%d\n", a)
	addTenToANumberUsingPointer(&a)
	fmt.Printf("%d\n", a)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(b)
	updateFirstIndex(b)
	fmt.Println(b)
}

func addTenToANumber(a int64) {
	a = a + 10
}

func addTenToANumberUsingPointer(a *int64) {
	(*a) = (*a) + 10
}

func updateFirstIndex(b []int) {
	b[0] = 0
}
