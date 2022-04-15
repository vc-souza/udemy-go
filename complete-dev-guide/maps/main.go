package main

import "fmt"

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}

func main() {
	// var countryCodes map[string]string

	// lastNames := make(map[string]string)

	// lastNames["Vinicius"] = "Carneiro"

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	// delete(colors, "greenw")

	// fmt.Printf("%v\n", countryCodes)
	// fmt.Printf("%v\n", lastNames)
	// fmt.Printf("%v\n", colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}
