package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // no more messages can be send in the channel
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
