package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("async ...")
	}()

	fmt.Println("sync ...")

	wg.Wait()
}
