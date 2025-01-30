package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<- c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might Be Down!")
		c <- "might be down from channel!"
		return 
	}

	fmt.Println(link, " Is Up!")
	c <- "is up!"
}
