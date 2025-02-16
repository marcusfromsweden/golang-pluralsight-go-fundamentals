// Package menu provides functionality for managing
// a menu of items and prices.
package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItems struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin)

// AddItem adds a new item to the menu
func AddItem() {
	fmt.Println("Enter item name:")
	itemName, _ := in.ReadString('\n')
	itemName = strings.TrimSpace(itemName)
	fmt.Println("New item name:", itemName)
	menu = append(menu, menuItems{name: itemName, prices: make(map[string]float64)})
}

// Print prints the menu to the console
func Print() {
	fmt.Println("Menu:")
	for _, item := range menu {
		fmt.Println(item.name)
		//fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s: %10.2f\n", size, price)
		}
	}
}
