package main

import (
	"fmt"
)

func main() {
	
	colors := map[string]string {
		"white": "#ffffff",
		"red":   "#ff0000",
		"green": "#blah",
	}

	iterateMap(colors)
}

func iterateMap(colors map[string]string) {
	for key, value := range(colors) {
		fmt.Println(key, value)
	}
}