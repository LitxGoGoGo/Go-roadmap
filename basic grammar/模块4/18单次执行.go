package main

import (
	"fmt"
	"sync"
	"time"
)

func Test18() {
	fmt.Println("函数执行")
}

func main1801() {
	var once sync.Once

	for i := 0; i < 10; i++ {
		go func() {
			//Test18(i)//调用函数,多次调用
			once.Do(Test18) //初始化一次可以直接使用,无需再初始化
		}()
	}

	time.Sleep(time.Second * 5)
}
