package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)
	arr = [3]string{"Coffee", "Orange juice", "Cordial"}
	fmt.Println(arr)
	fmt.Println(arr[0])
	arr[1] = "Tea"
	fmt.Println(arr[1])

	arr2 := arr
	arr2[2] = "Water"
	fmt.Println(arr, arr2)
}
