package main

import "fmt"

func main() {
	switch myVar := 9; true { // same as: switch myVar;
	case myVar < 10:
		fmt.Println("myVar is less than 10")
	case myVar == 10:
		fmt.Println("myVar is equal to 10")
	case myVar > 10:
		fmt.Println("myVar is greater than 10")
	}
}
