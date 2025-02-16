package main

import "fmt"

func main() {
	var someString string
	someString = "Hello, World!"
	fmt.Println(someString)

	var someInt int = 99
	fmt.Println(someInt)

	var someBool bool = true
	fmt.Println(someBool)

	someVar := 3.13
	fmt.Println(someVar)

	var convertedInt int = int(someVar)
	fmt.Println(convertedInt)
}
