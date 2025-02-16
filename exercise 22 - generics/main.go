package main

import (
	"fmt"
)

func main() {
	testScores := []int{90, 70, 50, 80, 60}
	floatTestScores := []float64{90.5, 70.5, 50.5, 80.5, 60.5}
	clonedTestScores := genericClone(testScores)
	clonedFloatTestScores := genericClone(floatTestScores)

	fmt.Println(&testScores[0], &clonedTestScores[0], clonedTestScores)
	fmt.Println(&floatTestScores[0], &clonedFloatTestScores[0], clonedFloatTestScores)

	mapOfTestScores := map[string]int{
		"John": 56,
		"Jane": 77,
	}
	floatMapOfTestScores := map[string]float32{
		"John": 56.5,
		"Jane": 77.5,
	}
	clonedMapOfTestScores := genericCloneMap(mapOfTestScores)
	clonedFloatMapOfTestScores := genericCloneMap(floatMapOfTestScores)

	fmt.Println(clonedMapOfTestScores)
	fmt.Println(clonedFloatMapOfTestScores)
}

func cloneMap(m map[string]int) map[string]int {
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func genericCloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func clone(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func genericClone[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	return result
}
