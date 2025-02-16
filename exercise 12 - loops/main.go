package main

import "fmt"

func main() {
	// infinite loop
	i := 1
	for {
		fmt.Println(i)
		i += 1
		if i > 4 {
			fmt.Println("Loop ends 1")
			break
		}
	}

	// loop to condition
	//for condition { ... }
	for i < 9 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Loop ends 2")

	// counter-based loop
	for i := 0; i < 6; i++ {
		fmt.Println(i)
	}
	fmt.Println("Loop ends 3")
}
