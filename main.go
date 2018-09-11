package main

import (
	"fmt"
	"io"
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

	// Takes the response body and writes it to the os.Stdout
	io.Copy(os.Stdout, resp.Body)
}
