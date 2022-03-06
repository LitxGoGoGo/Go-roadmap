package main

import "fmt"

//思想就是接口作为函数参数进行传递

type Humaner11 interface {
	SayHello11()
}
type Student11 struct {
	name  string
	score float64
}
type Teacher struct {
	name    string
	subject string
}

func (s *Student11) SayHello11() {
	fmt.Printf("大家好我是%s,我的成绩是%.2f\n", s.name, s.score)
}
func (t *Teacher) SayHello11() {
	fmt.Printf("大家好我是老师%s,我教的科目是%s", t.name, t.subject)
}

// WhoSayHello 将接口类型作为函数参数类型,实现多态
func WhoSayHello(human Humaner11) {
	human.SayHello11()
}
func main1101() {
	stu := Student11{
		name:  "Litx",
		score: 99,
	}

	tea := Teacher{
		name:    "爱因斯坦",
		subject: "物理",
	}

	/*
		var student Humaner11
		student = &stu //&取地址
		student.SayHello11()

		var teacher Humaner11
		teacher = &tea
		teacher.SayHello11()
	*/

	//以下是实现多态
	WhoSayHello(&stu)
	WhoSayHello(&tea)
}
