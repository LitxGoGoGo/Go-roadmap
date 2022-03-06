package main

import (
	"fmt"
	"sync"
)

func test() {
	fmt.Println("测试函数")
}

func main1301() {
	var pool sync.Pool
	//pool是一个协程安全的变量临时存储的数据格式
	pool.Put("Hello")
	pool.Put(true)
	pool.Put(123)
	pool.Put(test)

	v1 := pool.Get()
	fmt.Println(v1)
	v2 := pool.Get()
	fmt.Println(v2)
	v3 := pool.Get()
	fmt.Println(v3)
	v4 := pool.Get()
	fmt.Println(v4)
}

func main1302() {
	var pool sync.Pool
	pool.Put("Hello")
	v1 := pool.Get()
	fmt.Println(v1)

	pool.Put(123)
	v2 := pool.Get()
	fmt.Println(v2)

	pool.Put(true)
	v3 := pool.Get()
	fmt.Println(v3)
}
