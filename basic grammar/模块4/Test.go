package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func mainTest01() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendQuote(slice, "Hello")
	slice = strconv.AppendQuoteRune(slice, '我')
	slice = append(slice, '}')

	fmt.Println(string(slice))
}
func mainTest02() {
	a := make([]int, 10, 20)
	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
}
func mainTest03() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
func mainTest04() {
	// 声明一个等待组
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"https://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)
			// 通过参数传递url地址
		}(url)
	}
	// 等待所有的任务完成
	wg.Wait()
	fmt.Println("over")
}
