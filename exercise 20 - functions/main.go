package main

import "fmt"

func main() {
	name, otherName := "John", "Doe"
	fmt.Println("1: ", name, otherName)
	modifyValues(name, &otherName)
	fmt.Println("2: ", name, otherName)
	sum1 := sum(1, 2)
	fmt.Printf("Sum 1: %d\n", sum1)
	division, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Division 1: %d\n", division)
	}
	division2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Division 1: %d\n", division2)
	}

	sum2 := sumUsingNakedReturn(1, 2)
	fmt.Printf("Sum 2: %d\n", sum2)
}

func modifyValues(name string, otherName *string) {
	name = "Jane"
	*otherName = "Smith"
}

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func sumUsingNakedReturn(a, b int) (sum int) {
	sum = a + b
	return
}
