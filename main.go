package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// Create a slice of URLs
	URLs := []string{
		"https://google.com/",
		"https://linkedin.com/",
		"https://instagram.com/",
		"https://facebook.com/",
		"https://twitter.com/",
	}

	// Create a channel c
	c := make(chan string)

	// Launch goRoutines for each link
	for _, url := range URLs {
		go checkSite(url, c)
	}

	// Wait for result back from the channel.
	// Relaunch the same func in a gORoutine again, now with a 5s wait.
	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkSite(l, c)
		}(link)
	}
}

// Make an HTTP call to the site and print whether it works or not.
// Return the link in a channel
func checkSite(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is NOT working")
		c <- link
		return
	}
	fmt.Println(link, "works")
	c <- link
	return
}
