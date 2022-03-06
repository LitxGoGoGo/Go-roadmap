package main

import (
	"fmt"
	"time"
)

func main1201() {
	ch := make(chan int, 5)
	//超时判断
	exit := make(chan bool)

	go func() {
		for {
			select {
			case value := <-ch:
				fmt.Println(value)
			case <-time.After(time.Second * 5): //如果5s没有传递数据那么就执行连接超时操作
				fmt.Println("连接超时")
				exit <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		if i == 5 {
			time.Sleep(time.Second * 6)
		}
	}

	if <-exit {
		fmt.Println("数据销毁")
	}
	fmt.Println("程序结束")
}
