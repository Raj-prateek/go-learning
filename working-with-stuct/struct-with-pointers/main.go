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

type person struct {
	firstName string
	lastName  string
}

func main() {
	ram := person{firstName: "RAM", lastName: "GUPTA"}
	ramPointer := &ram
	ramPointer.updateFirstName("GRAM")
	ram.print()
	ram.updateFirstName("SHYAM")
	ram.print()
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
