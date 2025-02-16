package main

import (
	"fmt"
)

func main() {
	myArray := [3]string{"a", "b", "c"}
	for i, value := range myArray {
		fmt.Printf("%d: %s\n", i, value)
	}

	myMay := map[string]string{
		"Name": "May",
		"Age":  "20",
	}

	fmt.Println("Range over keys&values")
	for key, value := range myMay {
		fmt.Println(key, value)
	}

	fmt.Println("Range over keys")
	for key := range myMay {
		fmt.Println(key)
	}

	fmt.Println("Range over values")
	for _, value := range myMay {
		fmt.Println(value)
	}
}
