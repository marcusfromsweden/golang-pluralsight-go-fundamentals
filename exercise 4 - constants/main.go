package main

import "fmt"

func main() {
	const a = 42
	var someInt int = a
	var someFloat float64 = a //works since a is replaced by 42 before compilation
	fmt.Println(someInt, someFloat)

	const b float32 = 42.78
	//var someFloat2 float64 = b //not ok since b is declared as float32 and we are not converting it to float64
	var someFloat2 float64 = float64(b) //ok since we are converting b to float64
	fmt.Println(someFloat2)

	const (
		A2 = iota // 0
		B2        // 1
		C2        // 2
	)
	fmt.Println(A2, B2, C2)
	const (
		X = iota + 1  // 0 + 1 = 1
		Y = iota * 10 // 1 * 10 = 10
		Z = iota * 10 // 2 * 10 = 20
	)
	fmt.Println(X, Y, Z)
	const (
		_  = iota             // Skip 0
		KB = 1 << (iota * 10) // 1 << (1*10) = 1024
		MB = 1 << (iota * 10) // 1 << (2*10) = 1048576
		GB = 1 << (iota * 10) // 1 << (3*10) = 1073741824
	)
	fmt.Println(KB, MB, GB)
	const (
		_   = iota             // Skip 0
		KB2 = 1 << (iota * 10) // 1 << (1*10) = 1024
		MB2
		GB2
	)
	fmt.Println(KB2, MB2, GB2)
	const (
		A3, B3 = iota, iota + 10 // A = 0, B = 10
		C3, D3 = iota, iota + 10 // C = 1, D = 11
	)
	fmt.Println(A3, B3, C3, D3)

}
