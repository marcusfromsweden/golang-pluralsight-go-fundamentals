package main

import "fmt"

func main() {
	someInts := []int{1, 2, 3, 4, 5}
	someFloats := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	someStrings := []string{"a", "b", "c", "d", "e"}

	s1 := sum(someInts)
	s2 := sum(someFloats)
	s3 := sum(someStrings)

	fmt.Printf("Sum of someInts: %d\n", s1)
	fmt.Printf("Sum of someFloats: %f\n", s2)
	fmt.Printf("Sum of someStrings: %s\n", s3)
}

// addable is a type constraint that restricts the types that can be used with the sum function.
// all types supports += operator
type addable interface {
	int | float32 | string
}

func sum[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
