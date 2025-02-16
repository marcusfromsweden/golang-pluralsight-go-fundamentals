package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItems struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItems{
		{
			name:   "Coffee",
			prices: map[string]float64{"small": 5.99, "medium": 6.99, "large": 7.99},
		},
		{
			name:   "Tea",
			prices: map[string]float64{"small": 4.99, "medium": 5.99, "large": 6.99},
		},
		{
			name:   "Espresso",
			prices: map[string]float64{"single": 3.99, "double": 4.99},
		},
	}

	fmt.Println("Select an option and press enter:")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Println("Menu:")
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			//fmt.Printf("\t%#v: %#v\n", size, price)
			fmt.Printf("\t%10s: %10.2f\n", size, price)
		}
	}
}
