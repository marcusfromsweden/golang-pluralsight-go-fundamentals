package main

import (
	"bufio"
	"fmt"
	"main/menu"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

myLoop:
	for {
		choice := getUserInput()

		switch choice {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()
		case "q":
			break myLoop
		default:
			fmt.Println("Unknown option")
		}

	}
	fmt.Println("Goodbye!")
}

func getUserInput() string {
	fmt.Println("Select an option and press enter:")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add item")
	fmt.Println("q) Quit")
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)
	return choice
}
