package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)
	go func() {
		ch <- 42
	}()
	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}
