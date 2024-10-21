package main

import "fmt"

// func main() {
// 	colors := map[string]string{ // map is a structure in go that takes the keys of same type and the values of same type.
// 		"red":   "#ff0000",
// 		"green": "#fgh4344",
// 	}

// 	fmt.Println(colors)
// }

// --------------

// func main() {
// 	var colors map[string]string

// 	colors["red"] = "#ffrt453"

// 	fmt.Println(colors)
// }

// -----------------

// func main() {
// 	colors := make(map[string]string)

// 	colors["red"] = "#ffrt453"
// 	colors["green"] = "#sdf5455"

// 	delete(colors, "green")

// 	fmt.Println(colors)
// }

// ------------

func main() {
	colors := map[string]string{
		"red":   "redHex",
		"green": "greenHex",
		"white": "whiteHex",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hax := range c {
		fmt.Println("hax code for ", color, "is ", hax)
	}
}
