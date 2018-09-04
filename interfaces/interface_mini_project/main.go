package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://artifactory.deere.com/list/pypi-local/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}


	// To READ THE DATA
	// 1st ->
	/*
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	*/

	// 2nd ->
	/*
	// Make a byte slice and assign a predefined size for the byte slice
	bs := make([]byte, 9999)
	// This byte slice which gets passed, in the one which the Read method
	// add data to after reading from the Body.
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	*/

	// 3rd -> os.Stdout implements the writer interface and resp.Body implements the reader
	// io.Copy(os.Stdout, resp.Body)

	// 4th -> custom writer interface
	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs));

	return len(bs), nil
} 