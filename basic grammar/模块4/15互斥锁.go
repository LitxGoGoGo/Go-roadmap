package main

import (
	"fmt"
	"sync"
)

// Num 为全局变量进行赋值
var Num int

func Modify(wg *sync.WaitGroup, mu *sync.Mutex) {
	//给操作加锁
	mu.Lock()
	//延迟调用解锁
	defer mu.Unlock()
	Num++
	wg.Done()
}

func main1501() {
	var wg sync.WaitGroup
	//定义互斥锁
	var mu sync.Mutex
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go Modify(&wg, &mu)
	}

	wg.Wait()

	fmt.Println(Num)
}
