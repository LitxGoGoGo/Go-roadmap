package main

import (
	"errors"
	"fmt"
	"os"
)

//数组模拟队列
type Queue struct {
	maxSize int
	array   [5]int
	front   int //指向队列首部
	rear    int //指向尾部
}

//输入队列
func (que *Queue) AddQueue(Data int) (err error) {
	//先判断队列是否已经满
	if que.rear == que.maxSize-1 { //rear是队列尾部(包含最后元素)
		return errors.New("queue full")
	}

	que.rear++ //rear后移
	que.array[que.rear] = Data
	return
}

//输出队列
func (que *Queue) GetQueue() (val int, err error) {
	//判断队列是否为空
	if que.front == que.rear {
		return 1, errors.New("queue empty")
	}
	que.front++
	val = que.array[que.front]
	return val, err
}

//显示队列,找到队首遍历到队尾
func (que *Queue) ShowQueue() {
	//front不包含队首的元素
	for i := que.front + 1; i <= que.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, que.array[i])
	}
}

func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	for {
		fmt.Println("1. 输入add表示添加数据到队列")
		fmt.Println("2. 输入get表示从队列获取数据")
		fmt.Println("3. 输入show表示显示队列")
		fmt.Println("4. 输入exit表退出")

		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入你想要入队的数")
			var val int
			fmt.Scan(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
		fmt.Println()
	}

}
