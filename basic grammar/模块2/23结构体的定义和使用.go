package main

import "fmt"

//创建结构体
/*
type 结构体名称 struct {
	结构体成员 数据类型
	结构体成员 数据类型
}
*/

type student struct {
	id    int
	name  string
	age   uint8
	sex   rune
	score float64
	addr  string
}

func main2301() {
	//按照结构体顺序创建结构体变量
	stu := student{
		id:    1001,
		name:  "李天祥",
		age:   24,
		sex:   '男',
		score: 92,
		addr:  "广西玉林",
	}

	//可以先定义结构体变量再通过结构体成员进行赋值
	var stu1 student
	stu1.name = "李亮亮"
	stu1.age = 18
	stu1.id = 1002
	stu1.addr = "北京区"
	stu1.score = 87
	stu1.sex = '女'

	fmt.Printf("id:%d name:%s age:%d sex:%c score:%.2f addr:%s\n",
		stu.id, stu.name, stu.age, stu.sex, stu.score, stu.addr)

	fmt.Printf("id:%d name:%s age:%d sex:%c score:%.2f addr:%s\n",
		stu1.id, stu1.name, stu1.age, stu1.sex, stu1.score, stu1.addr)
}
