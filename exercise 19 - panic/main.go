package main

import "fmt"

func main() {
	a, b := 10, 2
	fmt.Println(divide(a, b))
	c, d := 10, 0
	fmt.Println(divide(c, d))
}

func divide(a, b int) int {
	//add defer recovery
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return a / b
}
