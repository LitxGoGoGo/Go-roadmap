package main

import "fmt"

type Student06 struct {
	id    int
	name  string
	score float64
}

type Person06 struct {
	id   int
	name string
}

func (stu *Student06) Modify06() {
	stu.name = "未命名"
	fmt.Println("我的名字改成了", stu.name)
}

func main0601() {
	stu := Student06{
		id:    1001,
		name:  "哈哈",
		score: 98,
	}
	stu.Modify06()
	//(&stu).Modify06() //go语言中将指针接收者优化成对象接收者了
	fmt.Println(stu)
}
