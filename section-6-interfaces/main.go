package main

import "fmt"

type bot interface {
    getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

// In this code we cannot use function overloading. as printGreeting has different parameters and still we cannot create 2 functions with the 
// same name
// also we want to reuse printGreeting function.
func main() {
	eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello There"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}