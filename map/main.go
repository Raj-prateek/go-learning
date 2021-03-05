package main

import "fmt"

/*
	In go map is a collection of key value pair.

	Map keys will be of same time as well as all
	it's value.

*/

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "000000",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	iterateColorsMap(colors)

	var sports map[string]string
	fmt.Println(sports)

	cars := make(map[string]string)
	cars["nios"] = "hyundai"
	cars["swift"] = "maruti"
	delete(cars, "swift")
	fmt.Println(cars)
}

func iterateColorsMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("%+v - %+v\n", color, hex)
	}
}
