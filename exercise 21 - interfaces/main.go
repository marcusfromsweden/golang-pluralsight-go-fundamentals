package main

import (
	"bytes"
	"fmt"
)

type printer interface {
	Print() string
}

type user struct {
	name string
	age  int
}

func (u user) Print() string {
	return fmt.Sprintf("Name: %s (%d)", u.name, u.age)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (m menuItem) Print() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n", m.name))
	for size, price := range m.prices {
		//buffer.WriteString(fmt.Sprintf("  %s: $%.2f\n", size, price))
		fmt.Fprintf(&buffer, "  %s: $%.2f\n", size, price)
	}
	return buffer.String()
}

func main() {
	var myPrinter printer
	myPrinter = user{"John Doe", 30}
	fmt.Println(myPrinter.Print())

	myPrinter = menuItem{
		name: "Pizza",
		prices: map[string]float64{
			"small":  5.99,
			"medium": 7.99,
			"large":  9.99,
		},
	}

	//myUser := myPrinter.(user) // failes with panic due to type assertion
	myUser, ok := myPrinter.(user)
	if !ok {
		fmt.Println("Type assertion (user) failed")
	}
	fmt.Println(myUser, ok)
	myMenuItem, ok := myPrinter.(menuItem)
	if !ok {
		fmt.Println("Type assertion (menuItem) failed")
	}
	fmt.Println(myMenuItem, ok)

	switch obj := myPrinter.(type) {
	case user:
		fmt.Println("User found:", obj)
	case menuItem:
		fmt.Println("MenuItem found:", obj)
	default:
		fmt.Println("Unknown type, but at least a printer ;)")
	}
}
