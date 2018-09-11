package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// saves the response and error of a get request to google.com
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// create a byte slice with lots of empty spaces, read function does not resize slice, only adds until full
	bs := make([]byte, 99999)

	// Saves the response body to the byte slice
	resp.Body.Read(bs)

	// prints out the body as a string
	fmt.Println(string(bs))
}
