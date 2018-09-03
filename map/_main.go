package main

import (
	"fmt"
)

func main() {
	// Ways of defining a map
	// 1st: below. But this creates nil map, entries cannot be added to it
	//var colors map[string]string

	// 2nd: below
	colors := map[string]string{
		"white": "#ffffff",
		"red":   "#ff0000",
		"green": "#blah",
	}
	// 3nd: below. Using built-in make command

	/*
		colors := make(map[string]string)
		colors["white"] = "#ffffff"

		// to delete: use delete keyword
		delete(colors, "white")
		fmt.Println(colors)
	*/

	iterateMap(colors)

}

func iterateMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
