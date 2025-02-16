package main

import "fmt"

func main() {
	switch myInt := 52; myInt {
	case 1: // if myInt is 1
		fmt.Println("Value is 1")
	case 2 + 3: // if myInt is 5
		fmt.Println("Value is 5")
	default: // if myInt is not 1 or 5
		fmt.Println("Value is not 1 or 5")
	}
}
