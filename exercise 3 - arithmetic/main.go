package main

import "fmt"

func main() {
	a, b := 5, 2
	fmt.Println("Addition: ", a+b)
	fmt.Println("Subtraction: ", a-b)
	fmt.Println("Multiplication: ", a*b)
	fmt.Println("Division: ", a/b)
	fmt.Println("Modulus: ", a%b)
	fmt.Println("convert to float: ", float64(a)/float64(b))

	fmt.Println("Equal to: ", a == b)
	fmt.Println("Greater than: ", a > b)
	fmt.Println("Less than: ", a < b)
	fmt.Println("Greater than or equal to: ", a >= b)
	fmt.Println("Less than or equal to: ", a <= b)
}
