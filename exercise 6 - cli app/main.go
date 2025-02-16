package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("what should I scream?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n') // using ` for "single variable"
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	fmt.Println(text)
}
