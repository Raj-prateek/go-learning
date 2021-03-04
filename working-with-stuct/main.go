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
	// First way for declaration struct
	shyam := person{"SHYAM", "KISHOR", contactInfo{"shyam@prateek.com", "99999999"}}
	fmt.Println(shyam)

	// Second way for declaration struct
	ram := person{firstName: "RAM", lastName: "RAJ", contactinfo: contactInfo{"ram@prateek.com", "99999999"}}
	fmt.Println(ram)

	// Third way for declaration struct
	/*
		In this approach `raj` person propties will be assigned a value of empty string.

		As per the properties data type the default varies as defined below:
		TYPE   | Value
		______________
		string | ""
		int    | 0
		float  | 0
		bool   | false
	*/
	var raj person
	fmt.Println(raj)
	fmt.Printf("%+v\n", raj)
	raj.firstName = "RAJ"
	raj.lastName = "KISHOR"
	raj.contactinfo = contactInfo{"ram@prateek.com", "99999999"}
	fmt.Println(raj)

}
