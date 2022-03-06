package main

import (
	"fmt"
	"time"
)

func main0701() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	//单向channel
	var send chan<- int = ch    //单向发送channel
	var receive <-chan int = ch //单向接收channel

	send <- 123 //ok
	value := <-receive
	fmt.Println(value)
	time.Sleep(time.Second * 5)
}
