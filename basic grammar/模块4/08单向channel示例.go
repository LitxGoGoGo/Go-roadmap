package main

import (
	"fmt"
	"time"
)

//单向channel生产者,只读
func producer(in <-chan int) {
	for value := range in {
		fmt.Println(value)
	}
}

//单向channel消费者,只写
func consumer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

func main0801() {
	ch := make(chan int, 3)
	go consumer(ch)
	go producer(ch)
	time.Sleep(time.Second)
}
