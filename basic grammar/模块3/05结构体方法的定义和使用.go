package main

import "fmt"

type Student05 struct {
	id    int
	name  string
	score float64
}

//给结构体定义方法
func (s Student05) sayMyInfo() {
	fmt.Println("我的名字是", s.name, ",学号是", s.id, "成绩麻麻地,也就", s.score, "分")
}

func main0501() {
	var stu = Student05{
		id:    1001,
		name:  "李天祥",
		score: 98.45,
	}
	stu.sayMyInfo()
	fmt.Println(stu)
}
