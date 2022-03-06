package main

import (
	"fmt"
	"time"
)

func main0501() {
	//make(chan 数据类型, 容量)
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("创建channel的数据", len(ch), cap(ch))
		}
	}()
	time.Sleep(time.Second * 5)

	for i := 0; i < 10; i++ {
		fmt.Println("使用channel的数据", len(ch), cap(ch))
		value := <-ch
		fmt.Println(value)
		time.Sleep(time.Second * 5)
	}

}
