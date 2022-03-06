package main

import (
	"fmt"
	"sync"
	"time"
)

func Test17(wg *sync.WaitGroup, cond *sync.Cond, i int) {
	cond.L.Lock()
	fmt.Println("协程", i, "已经加锁")

	cond.Wait()
	fmt.Println("协程", i, "已经解锁")

	cond.L.Unlock()

}

func main1701() {
	var wg sync.WaitGroup
	var cond *sync.Cond
	cond = sync.NewCond(new(sync.Mutex))
	//创建等待组的个数
	//wg.Add(10)
	for i := 0; i < 10; i++ {
		go Test17(&wg, cond, i)
	}
	//发送信号
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		cond.Signal()
	}

	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("所有信号已经发送,同步信号已经收到")
}
