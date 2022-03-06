package main

import "fmt"

type Student03 struct {
	Person
	subject string
	score   float64
}

type Person struct {
	Object
	age uint8
	sex string
}

type Object struct {
	id   int
	name string
}

func main0301() {
	var stu Student03 = Student03{
		Person: Person{
			Object: Object{
				id:   1001,
				name: "李天祥",
			},
			age: 26,
			sex: "男",
		},
		subject: "数学",
		score:   96,
	}

	//可以对变量进行修改
	stu.name = "李牛牛"
	fmt.Println(stu.name, stu.id)
	fmt.Println(stu)
}
