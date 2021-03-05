package main

import "fmt"

type contactInfo struct {
	email    string
	mobileno string
}

type person struct {
	firstName   string
	lastName    string
	contactinfo contactInfo
}

func main() {
	ram := person{firstName: "RAM", lastName: "RAJ", contactinfo: contactInfo{"ram@prateek.com", "99999999"}}
	ram.print()
	ram.setFirstName("Rajat")
	ram.print()
}

func (p person) setFirstName(firstName string) {
	p.firstName = firstName
}

func (p person) setLastName(lastName string) {
	p.lastName = lastName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
