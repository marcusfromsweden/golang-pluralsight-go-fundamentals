package main

import (
	"fmt"
	"slices"
)

func main() {
	var mySlice []string
	fmt.Println(mySlice)
	mySlice = []string{"Coffee", "Orange juice", "Cordial"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])
	mySlice[1] = "Tea"
	fmt.Println(mySlice[1])

	mySlice = append(mySlice, "Soda")

	otherSlice := mySlice
	otherSlice[2] = "Water"
	fmt.Println(mySlice, otherSlice) // Cordial will be replaced by Water in both slices

	mySlice = slices.Delete(mySlice, 1, 2)
	fmt.Println(mySlice) // Coffee, Water, Soda

	myMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(myMap)
}
