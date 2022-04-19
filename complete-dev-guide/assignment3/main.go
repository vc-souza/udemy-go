package main

import (
	"fmt"
	"io"
	"os"
)

func printFileError(filepath string, err error) {
	if err == nil {
		return
	}

	fmt.Printf("Could not open file '%v': %v\n", filepath, err)
	os.Exit(1)
}

func fileToConsole(filepath string) {
	fmt.Printf("### File '%v':\n", filepath)

	f, oErr := os.Open(filepath)
	printFileError(filepath, oErr)

	_, coErr := io.Copy(os.Stdout, f)
	printFileError(filepath, coErr)

	cErr := f.Close()
	printFileError(filepath, cErr)

}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No filepath argument was given")
		os.Exit(1)
	}

	for _, filepath := range os.Args[1:] {
		fileToConsole(filepath)
	}
}
