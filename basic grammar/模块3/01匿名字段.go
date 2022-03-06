package main

import (
	"fmt"
)

/*
在go中对于继承关系通过匿名字段来实现
*/

type person01 struct {
	name string
	age  uint8
	sex  string
}

/*
student能继承person的属性
*/
type student01 struct {
	person01
	grade string
	score float64
}

type student02 struct {
	per   person01
	grade string
	score float64
}

func main0101() {
	var stu = student01{
		person01: person01{
			name: "李天祥",
			age:  24,
			sex:  "男",
		},
		grade: "大四",
		score: 92,
	}
	fmt.Println(stu)
	fmt.Println(stu.name, stu.age, stu.sex, stu.grade, stu.score)
}

func main0102() {
	var stu = student02{
		per: person01{
			name: "李牛牛",
			age:  25,
			sex:  "男",
		},
		grade: "大三",
		score: 95,
	}
	fmt.Println(stu.per, stu.per.sex, stu)
}
