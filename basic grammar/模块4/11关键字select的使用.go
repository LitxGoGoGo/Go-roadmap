package main

import (
	"fmt"
	"time"
)

func main1101() {
	/*
		select{
		case <-chan:
		  //当channel成功读取数据执行当前的分支
		case chan<-1:
		  //当channel写入数据 执行当前的分支
		case chan<-2:
		  //当channel写入数据 执行当前的分支
		default :
		  //如果以上都没有触发则执行默认操作
		}
	*/
	ch := make(chan int)
	go func() {
		select {
		case value := <-ch:
			fmt.Println("ch成功读取数据")
			fmt.Println(value)
		case ch <- 1:
			fmt.Println("ch成功写入数据")
		}
	}()
	time.Sleep(time.Second * 1)
	ch <- 1
	time.Sleep(time.Second * 1)
}

func Fibonacci(ch chan int, exit chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-exit:
			fmt.Println("完成")
			return
		}
	}
}

func main1102() {
	ch := make(chan int)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 30; i++ {
			value := <-ch
			fmt.Printf("%d ", value)
		}
		exit <- true
	}()
	Fibonacci(ch, exit)

	time.Sleep(time.Second * 2)
}
