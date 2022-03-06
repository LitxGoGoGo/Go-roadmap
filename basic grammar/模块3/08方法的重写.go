package main

import "fmt"

type Person08 struct {
	name string
	sex  string
}
type Student08 struct {
	Person08
	subject string
	score   float64
}

func (per Person08) sayHi() {
	fmt.Println("我是父类的方法sayHi")
}

func (stu Student08) sayHi() {
	fmt.Println("重写了方法")
}
func main0801() {
	stu := Student08{
		Person08: Person08{
			name: "崔梦洋",
			sex:  "女",
		},
		subject: "以太坊前端工程师",
		score:   59.9,
	}
	//默认采用就近原则 使用本对象的方法
	//stu.SayHi()
	stu.Person08.sayHi()
	stu.sayHi() //就近原则找最近的那一个
}
