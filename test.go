package main

import (
	"fmt"
)

func main() {
	var (
		ch   = make(chan int, 5)
		done = make(chan struct{})
	)
	go sender(ch)
	go receiver(ch, done)
	<-done
}

func sender(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("wrote to channel", i)
	}
	close(ch)
}

func receiver(ch <-chan int, done chan<- struct{}) {
	for num := range ch { // итерация по каналу
		fmt.Println("read from channel", num)
	}
	close(done)
}
