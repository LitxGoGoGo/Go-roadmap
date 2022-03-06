package main

import "fmt"

type Person07 struct {
	name string
	age  uint8
}

type Student07 struct {
	Person07
	subject string
	score   float64
}

// SayPerson 给父类创建方法
func (per Person07) SayPerson() {
	fmt.Printf("大家好我是%s，我今年%d岁\n", per.name, per.age)
}

// SayStuedent 给子类创建方法
func (stu Student07) SayStuedent() {
	fmt.Println("我是子类的方法hi")
}
func main0701() {
	per := Person07{
		name: "李天祥",
		age:  25,
	}

	per.SayPerson()

	stu := Student07{
		Person07: Person07{
			name: "李牛牛",
			age:  27,
		},
		subject: "前端",
		score:   98,
	}

	stu.SayStuedent() //子类自己定义的方法
	stu.SayPerson()   //在父类定义的方法
}
