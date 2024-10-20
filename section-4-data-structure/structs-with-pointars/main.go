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
		iqbalPointer := &iqbal // address of the value for iqbal in Hexadecimal
		iqbalPointer.updateName("sabbir")

		writing the above lines are equivalent to:
		iqbal.updateName("sabbir")

		if you use pointer to structs in the reciever of a function
	*/
	iqbal.updateName("sabbir") // if called with a type and the functions reciever parameter is of type pointer to a type then go by default gets the address of the value and calls the function with the pointer
	iqbal.print()
}

// in here *person is a completely different type from person type
func (pointerToPerson *person) updateName(newFirstName string) { // pointerToPerson is of type (pointer to a person)
	(*pointerToPerson).firstName = newFirstName // *pointerToPerson means the value in the address of ponterToPerson in memory
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
