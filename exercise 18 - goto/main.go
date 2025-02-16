package main

import "fmt"

func main() {
	x := 5

	if x < 10 {
		fmt.Println("X is small")
		goto EXIT
	}

	fmt.Println("X is large")

EXIT:
	fmt.Println("Exiting program")
}
