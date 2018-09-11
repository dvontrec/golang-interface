package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// saves the response and error of a get request to google.com
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Creates a new logWriter
	lw := logWriter{}

	// Takes the response body and writes it to the os.Stdout
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	// prints the byte slice as a string
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes: ", len(bs))
	// returns the number of bytes and passes no error
	return len(bs), nil
}
