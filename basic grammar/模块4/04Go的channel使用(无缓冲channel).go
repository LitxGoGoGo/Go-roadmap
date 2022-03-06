package main

import "fmt"

func main0401() {
	//无缓冲的channel定义
	//用make进行定义
	//var ch chan int = make(chan int)

	//自动推导类型进行定义
	ch := make(chan int)
	//fmt.Println(ch)
	//fmt.Printf("%T", ch)
	//fmt.Println(len(ch), cap(ch))
	go func() {
		for i := 1; i < 5; i++ {
			ch <- i
		}
	}()

	for i := 1; i < 5; i++ {
		value := <-ch
		fmt.Println(value)
	}
}
