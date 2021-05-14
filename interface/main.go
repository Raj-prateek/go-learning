package main

import "fmt"

/*
	1. Interfaces are not generic type.
	2. Interfaces are 'implict'.
	3. Interfaces are a contract to help us manage types.
*/
type bot interface {
	//We can also define argument which is need to be passed in the definition 
	getGreeting() string
}

type englishBot struct {}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}


	printGreeting(eb)
	printGreeting(hb)
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi, There!"
}

func (hindiBot) getGreeting() string {
	return "Namaste!"
}
