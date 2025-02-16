package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	dividend1, divisor1 := 10, 2
	fmt.Printf("Dividend: %d\n", divide(dividend1, divisor1))

	dividend2, divisor2 := 10, 0
	fmt.Printf("Dividend: %d\n", divide(dividend2, divisor2))
}

func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("Error: %v\n", msg)
		}
	}()
	return dividend / divisor
}
