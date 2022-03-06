package main

import (
	"errors"
	"fmt"
	"os"
)

//定义结构体来管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int //指向队首
	tail    int //指向队尾
}

//添加入队操作
func (cil *CircleQueue) Push(val int) (err error) {
	if cil.IsFull() {
		return errors.New("队列已满")
	}
	cil.array[cil.tail] = val
	cil.tail = (cil.tail + 1) % cil.maxSize
	return
}

//出队操作
func (cil *CircleQueue) Pop() (val int, err error) {
	if cil.IsEmpty() {
		return 0, errors.New("队列已经为空")
	}
	//取出 head指向队首,并且包含队首元素
	val = cil.array[cil.head]
	cil.head = (cil.head + 1) % cil.maxSize
	return
}

//判断环形队列是否为满
func (cil *CircleQueue) IsFull() bool {
	return (cil.tail+1)%cil.maxSize == cil.head
}

//环形队列是否为空
func (cil *CircleQueue) IsEmpty() bool {
	return cil.head == cil.tail
}

//取出环形队列有多少个元素
func (cil *CircleQueue) Size() int {
	return (cil.tail + cil.maxSize - cil.head) % cil.maxSize
}

//显示队列Show
func (cil *CircleQueue) ListQueue() {
	//取出当前队列有多少元素
	size := cil.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量指向head
	tempHead := cil.head
	for i := 0; i < size; i++ {
		fmt.Printf("arry[%d]=%d\t", tempHead, cil.array[tempHead])
		tempHead = (tempHead + 1) % cil.maxSize
	}
	fmt.Println()
}
func main() {
	//先创建一个环形队列
	circleQueue := CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add表示添加数据到队列")
		fmt.Println("2. 输入get表示从队列获取数据")
		fmt.Println("3. 输入show表示显示队列")
		fmt.Println("4. 输入exit表退出")

		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入你想要入队的数")
			fmt.Scan(&val)
			err := circleQueue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := circleQueue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数", val)
			}
		case "show":
			circleQueue.ListQueue()
		case "exit":
			os.Exit(0)
		}
		fmt.Println()
	}

}
