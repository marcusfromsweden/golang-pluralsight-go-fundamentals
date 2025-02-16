package main

import "fmt"

func main() {
	var myMap map[string][]string
	fmt.Println(myMap)

	myMap = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappicino"},
		"tea":    {"Tea", "Green Tea", "Black Tea"},
	}
	fmt.Println(myMap)

	fmt.Println(myMap["coffee"])

	myMap["other"] = []string{"Hot chocolate", "Milk"}
	fmt.Println(myMap)

	delete(myMap, "tea")
	fmt.Println(myMap)
	value, ok := myMap["tea"]
	fmt.Println(value, ok)

	otherMap := myMap
	otherMap["coffee"] = []string{"Coffee"}
	myMap["other"] = []string{"Milk"}
	fmt.Println(myMap)
	fmt.Println(otherMap)
}
