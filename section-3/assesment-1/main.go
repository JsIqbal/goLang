package main

import "fmt"

type numbers []int

func main() {
	numbers := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers.evenOrOdd()

}

func (n numbers) evenOrOdd() {
	for _, num := range n {
		if num%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

/*
Create a new package
In the main funciton, create a slice of ints from 0 through 10

Iterate through the slice. For each element in the slice check to see if the number is even or odd

If the value is even, print out "even". if it is odd, print out "odd"
*/
