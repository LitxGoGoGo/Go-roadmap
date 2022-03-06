package main

import "fmt"

//接口理解作为一种行为规范
/*
如要满足接口定义接口内的方法必须全部实现,只能多不能少
type 接口名er interface{
	方法名(参数) 返回值
}
*/

// Humaner 定义接口
type Humaner interface {
	SayHello() //没有func关键字
	GetAge() uint8
	GetScore() float64
}

// Student 定义结构体
type Student10 struct {
	name  string
	age   uint8
	score float64
}

//使用方法将结构体和接口连接起来
func (s Student10) SayHello() {
	fmt.Printf("大家好,我是%s,我今年%d岁了,我的成绩是%.2f\n", s.name, s.age, s.score)
}
func (s Student10) GetAge() uint8 {
	return s.age
}
func (s Student10) GetScore() float64 {
	return s.score
}

func main1001() {
	stu := Student10{
		name:  "Litx",
		age:   25,
		score: 96,
	}

	/*	stu.SayHello()
		stu.GetAge() 未通过接口进行方法的实现 */

	//接口的使用
	var human Humaner
	//如果对象是值的对象接收者可以直接将对象stu赋值给接口类型
	human = stu
	human.SayHello()
	human.GetAge()
	human.GetScore()
}
