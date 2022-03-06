package main

import "fmt"

type Object04 struct {
	id   int
	name string
}

type Person04 struct {
	Object04
	age uint8
	sex string
}

type Student04 struct {
	//Object04
	Person04
	subject string
	score   float64
}

func main0401() {
	stu := Student04{
		//Object04: Object04{
		//	id:   1001,
		//	name: "李obj",
		//},
		Person04: Person04{
			Object04: Object04{
				id:   1002,
				name: "李per",
			},
			age: 24,
			sex: "男",
		},
		subject: "前端",
		score:   93,
	}
	fmt.Println(stu)
	//对于结构体内置的变量查找查到第一个就返回
	fmt.Println(stu.name, stu.Person04.name)
}
