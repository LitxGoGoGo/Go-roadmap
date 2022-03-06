package main

import (
	"fmt"
	"sync"
	"time"
)

func Demo1(wg *sync.WaitGroup) {
	fmt.Println("Demo1 函数执行完成")
	//当前组成员完成任务
	wg.Done()
}
func Demo2(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 3)
	fmt.Println("Demo2 函数执行完成")
	wg.Done()
}
func main1401() {
	var wg sync.WaitGroup
	//添加组成员
	wg.Add(2)

	go Demo1(&wg)
	go Demo2(&wg)
	//阻塞式的等待  只有等待组的计数器归零 才继续执行
	wg.Wait()
	fmt.Println("继续执行后续代码")
}
