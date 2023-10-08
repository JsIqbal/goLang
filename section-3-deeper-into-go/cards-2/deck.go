package main

import "fmt"

// we will create a custom type named deck which will be a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
