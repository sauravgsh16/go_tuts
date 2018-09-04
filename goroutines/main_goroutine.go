package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://artifactory.deere.com/list/pypi-local/",
		"http://python.org",
	}

	// channel of type chan which expects type string data to be transferred
	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	// Receiving a message over a Channel is blocking operation in nature
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link + "Might be down"
		os.Exit(1)
	}
	c <- link + " is up!!"
}