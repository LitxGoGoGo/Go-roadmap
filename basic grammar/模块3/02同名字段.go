package main

import "fmt"

type Person02 struct {
	name string
	age  uint8
	sex  string
}
type Student02 struct {
	Person02
	name    string
	subject string
	score   float64
}

func main0201() {
	stu := Student02{
		Person02: Person02{
			name: "DI第",
			age:  29,
			sex:  "26",
		},
		name:    "didi",
		subject: "中文",
		score:   98,
	}

	stu.name = "滴滴"
	stu.Person02.name = stu.name

	fmt.Printf("%+v\n", stu)
}
