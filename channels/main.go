package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com", 
		"http://facebook.com",
		"http://twitter.com", 
		"http://linkedin.com",
		"http://amazon.com",
	}

	c := make(chan string) // Created a brand new channel

	for _, w := range links {
		go checkLink(w, c) // Go routine call
	}
	
	/* 
		# Pull response from channel

		fmt.Println(<-c)
	*/

	/* 
		# Infinite checking status of the websites
	
		for {
			go checkLink(<-c, c)
		}

		for l := range c {
			go checkLink(l, c)
		}
	*/

	for l := range c {
		// Literal Functions
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link,c)
		}(l)
	}
}

func checkLink (w string, c chan string) {
	_, err := http.Get(w)
	if err != nil {
		fmt.Println(w, "might be down !")
		c <- w
		return
	}
	fmt.Println(w, "OK!")
	c <- w // Push value into the channel
}
