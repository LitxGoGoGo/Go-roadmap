package main

import (
	"fmt"
	"time"
)

func main0901() {
	//time.Timer{}延时器所在的对象
	//延迟调用
	<-time.After(time.Second * 2)
	fmt.Println("两秒已到1")

	time.Sleep(time.Second * 2)
	fmt.Println("两秒已到2")

	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("两秒已到3")
}

func main0902() {
	//定时器的停止
	timer := time.NewTimer(time.Second * 2)
	go func() {
		fmt.Println("协程开始执行")
		<-timer.C
		fmt.Println("协程执行完成")
	}()
	//计时器停止
	timer.Stop()

	time.Sleep(time.Second * 6)
}

func main0903() {
	//定时器的重置
	timer := time.NewTimer(time.Second * 60)
	ok := timer.Reset(time.Second * 5)
	if ok {
		fmt.Println("计时器已经重置")
	} else {
		fmt.Println("计时重置失败")
	}
	<-timer.C
	fmt.Println("时间到")
}
