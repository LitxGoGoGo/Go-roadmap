package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for i := 0; i < 6; i++ {
		fmt.Println("被调函数", i)
		time.Sleep(time.Second)
	}
}

func main0101() {
	go NewTask()

	for i := 0; i < 6; i++ {
		fmt.Println("主调函数", i)
		time.Sleep(time.Second)
	}
}
