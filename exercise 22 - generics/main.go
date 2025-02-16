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
	testScores := []int{90, 70, 50, 80, 60}
	c := clone(testScores)
	fmt.Println(&testScores[0], &c[0], c)
}

func clone(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
