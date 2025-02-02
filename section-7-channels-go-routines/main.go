// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	links := []string{
// 		"http://google.com",
// 		"http://youtube.com",
// 		"http://facebook.com",
// 	}

// 	c := make(chan string)

// 	for _, link := range links {
// 		go checkLink(link, c)
// 	}

// 	fmt.Println(<- c)
// 	fmt.Println(<- c)
// 	fmt.Println(<- c)
// }

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, " Might Be Down!")
// 		c <- "might be down from channel!"
// 		return 
// 	}

// 	fmt.Println(link, " Is Up!")
// 	c <- "is up!"
// }


package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://youtube.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i ++ {
		fmt.Println(<-c)
	}
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
