package main

import (
	"time"
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://artifactory.deere.com/list/pypi-local/",
		"http://python.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	// The below for loop syntax creates an infinite loop
	/*
	for {
		go checkStatus(<-c, c)
	}
	*/

	// An proper alternative to loop is as:
	// This declaration means, we are going to run an iteration, when we receive a message on the channel c
	for l := range c {
		// We want to add a 5 sec delay between to calls.
		// 1) If we place the delay in the main goroutine, the execution will wait for 5 sec.
		// We don't want that because there might be instances that someother goroutine completes
		// and returns data to the channel in the main goroutine, in such case, the main goroutine
		// will not process it till it wakes up after the sleep. Calls will start stacking in such scenario
		// 2) We can place the sleep in the "checkStatus" function and it will work, but we don't want that
		// Not good coding practise since, whenever we call checkstatus, we expect the function to return immediately.

		// Thus in order to solve above concerns, we make use of function literals.
		// Function literals are same as lambda in python
		// We just need to invoke the function literal after definition with '()'

		// The reason why we passed l as a function argument, please please please please look at the last lecture again
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkStatus(link, c)
		}(l)
	}

}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up")
	c <- link
}