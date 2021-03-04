package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, x := range numbers {
		switch x % 2 {
		case 1:
			fmt.Printf("%d is odd\n", x)
			break
		case 0:
			fmt.Printf("%d is even\n", x)
			break
		}
	}
}
