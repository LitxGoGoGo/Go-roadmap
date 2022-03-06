package main

import "fmt"

type Human13 interface {
	SayHi13()
}

type Personer13 interface {
	Human13 //继承了Human13接口
	SayHello()
}

type Person13 struct {
	name string
	age  uint8
	sex  string
} //没有实现SayHello

type Student13 struct {
	Person13
	score float64
}

func (p *Person13) SayHi13() {
	fmt.Println("Person13 SayHi13")
}

func (s *Student13) SayHi13() {
	fmt.Println("Student SayHi")
}

//func (p *Person13) SayHello() {
//	fmt.Println("Person SayHello")
//}

func (s *Student13) SayHello() {
	fmt.Println("Student SayHello")
}

func main() {
	stu := Student13{
		Person13: Person13{
			name: "Litx",
			age:  24,
			sex:  "男",
		},
		score: 98,
	}

	//per := Person13{
	//	name: "Hct",
	//	age:  28,
	//	sex:  "女",
	//}

	var human Human13     //子集
	var person Personer13 //超集

	person = &stu
	//将超集赋值给子集,可以将超集赋值给子集,因为超集有子集所有实现的方法
	human = person
	human.SayHi13()
}
