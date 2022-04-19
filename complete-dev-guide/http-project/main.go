package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logger struct{}

func (logger) Write(p []byte) (int, error) {
	return fmt.Print(string(p))
}

func main() {
	response, getErr := http.Get("https://google.com")

	if getErr != nil {
		fmt.Printf("Error: %v\n", getErr)
		os.Exit(1)
	}

	// Low-level, ugly, evil!
	// rawBody := make([]byte, 999999999)
	// response.Body.Read(rawBody)
	// fmt.Printf("Response: %v\n", string(rawBody))

	// Ready to go Writer (File)
	// io.Copy(os.Stdout, response.Body)

	io.Copy(logger{}, response.Body)
}
