package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	myError := errors.New("This is an error")
	fmt.Println(myError)

	myError2 := fmt.Errorf("This error wraps the first one: %w", myError)
	fmt.Println(myError2)

	return

	f, err := os.Open("filename.ext")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read from the file until new line
	b := make([]byte, 100)
	n, err := f.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b[:n]))
}
