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

type menu []menuItems

func (m menu) print() {
	fmt.Println("Menu:")
	for _, item := range m {
		fmt.Println(item.name)
		for size, price := range item.prices {
			fmt.Printf("\t%10s: %10.2f\n", size, price)
		}
	}
}

func (m *menu) add() {
	fmt.Println("Enter item name:")
	itemName, _ := in.ReadString('\n')
	itemName = strings.TrimSpace(itemName)
	*m = append(data, menuItems{name: itemName, prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin)

// AddItem adds a new item to the menu
func AddItem() {
	data.add()
}

// Print prints the menu to the console
func Print() {
	data.print()
}
