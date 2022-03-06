package main

import (
	"fmt"
	"runtime"
	"time"
)

func main0301() {

	go func() {
		//通过defer延迟调用函数
		defer fmt.Println("defer延迟调用")
		func() {
			fmt.Println("匿名函数的调用")
			runtime.Goexit()
			fmt.Println("匿名函数调用结束")
		}()
	}()

	//延迟main函数的结束
	time.Sleep(time.Second * 10)
}
