package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "bye"
}

// In here as slice and array uses pass by referrence that is why we do not need to use a pointer to work with any addresses
