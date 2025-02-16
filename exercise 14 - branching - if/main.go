package main

import "fmt"

func main() {
	myInt := 5
	if myInt < 10 {
		fmt.Println("myInt is less than 10")
	} else if myInt == 10 {
		fmt.Println("myInt is equal to 10")
	} else {
		fmt.Println("myInt is greater than or equal to 11")
	}
	fmt.Println("if 1 done")

	if otherInt := 10; otherInt < 11 {
		fmt.Println("otherInt is less than 11")
	}
	fmt.Println("if 2 done")
}
