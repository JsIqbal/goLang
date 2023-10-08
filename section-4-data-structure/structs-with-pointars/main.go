package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	number string
	zip    int64
}

func main() {
	iqbal := person{
		firstName: "Iqbal",
		lastName:  "Hossain",

		contactInfo: contactInfo{
			number: "8801403229479",
			zip:    1230,
		},
	}

	/*
		iqbalPointer := &iqbal
		iqbalPointer.updateName("sabbir")

		writing the above lines are equivalent to:
		iqbal.updateName("sabbir")

		if you use pointer to structs in the reciever of a function
	*/
	iqbal.updateName("sabbir")
	iqbal.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
