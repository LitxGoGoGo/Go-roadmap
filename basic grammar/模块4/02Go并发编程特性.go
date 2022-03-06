package main

import (
	"fmt"
	"time"
)

func Test() {
	for i := 0; i < 5; i++ {
		fmt.Println("被调函数", i)
		time.Sleep(time.Second)
	}
}

func main0201() {
	go Test()
	fmt.Println("主函数执行完成")
}
