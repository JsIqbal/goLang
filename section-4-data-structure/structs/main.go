package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	number string
	zip    int64
}

func main() {
	// iqbal := person{"Iqbal", "Hossain"} // The object assumes Iqbal is the firstname and hossain in the lastname
	iqbal := person{ // another way to define embedded structs
		firstName: "Iqbal",
		lastName:  "Hossain",

		// contact: contactInfo{ // one way to use embeded structs
		contactInfo: contactInfo{ // if we have key and value same in the struct then we can use only contactInfo
			number: "8801403229479",
			zip:    1230,
		},
	}

	iqbal.updateName("sabbir") // this will not work as go is a pass by value language. so the struct will never be updated in this manner.
	iqbal.print()
}

func (p person) updateName(newFirstName string) { // by default go is a pass by value type language. so in here go is creating a compy of the struct. saving it into a different location in memory. and modifying that memory
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
