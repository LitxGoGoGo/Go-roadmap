package main

import (
	"fmt"
)

func main0601() {
	//创建无缓存的channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	//遍历通道数据的两种方法

	//for {
	//	if value, ok := <-ch; ok {
	//		fmt.Println(value)
	//	} else {
	//		break
	//	}
	//}

	for value := range ch {
		fmt.Println(value)
	}
}
