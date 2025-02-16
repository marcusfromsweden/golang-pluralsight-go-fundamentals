package main

import "time"

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	/* 	go func() { ch1 <- "message to ch1" }()

	   	go func() { ch2 <- "message to ch2" }()
	*/
	time.Sleep(10 * time.Millisecond)

	// select picks the first channel that is ready and receives from it (or sends to it).
	// will be random
	select {
	case msg := <-ch1:
		println("from ch1", msg)
	case msg := <-ch2:
		println("from ch2", msg)
	default:
		println("default")
	}

}
