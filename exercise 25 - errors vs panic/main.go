package main

import (
	"errors"
	"fmt"
)

func main() {
	/* 	fmt.Println(divide(10, 2))
	   	result, err := divideWithError(10, 0)
	   	if err != nil {
	   		fmt.Println("Error:", err)
	   		return
	   	}
	   	fmt.Println(result) */

	result2, err2 := divide2(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println(result2)
}

func divide(a, b int) int {
	return a / b
}

func divideWithError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func divide2(a, b int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return a / b, nil
}
