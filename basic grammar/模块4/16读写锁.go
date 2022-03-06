package main

import (
	"fmt"
	"sync"
	"time"
)

//定义全局变量
var Value int

//为全局部变量进行赋值
func write(rwmu *sync.RWMutex, i int) {
	//写入锁
	rwmu.Lock()
	defer rwmu.Unlock()
	Value = i
	fmt.Print("W:", Value, " ")
}

//读取全局变量内容
func read(rwmu *sync.RWMutex) {
	//读取锁,如果当前数据添加了读取锁,其他协程可以读取但是不能写入
	rwmu.RLock()
	defer rwmu.RUnlock()
	result := Value
	fmt.Print("R:", result, " ")
}
func main1601() {
	//定义读写锁
	var rwmu sync.RWMutex

	//多次的写入和读取,先写入后读取
	for i := 0; i < 100; i++ {
		go write(&rwmu, i)
		go read(&rwmu)
	}
	fmt.Println()
	time.Sleep(time.Second * 5)
}
