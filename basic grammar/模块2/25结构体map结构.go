package main

import "fmt"

//定义结构体
type student25 struct {
	name  string
	age   uint8
	sex   rune
	score float64
	addr  string
}

func main2501() {
	//创建map,对应关系为key为学号,value为学生信息数据类型为student25
	stus := make(map[int]student25)
	stus[1001] = student25{
		name:  "刘能",
		age:   52,
		sex:   '男',
		score: 88,
		addr:  "象牙山51号",
	}
	stus[1002] = student25{
		"谢广坤",
		51,
		'男',
		89,
		"广东东莞",
	}
	stus[1098] = student25{
		name:  "王老七",
		age:   55,
		sex:   '男',
		score: 99.5,
		addr:  "象牙山1号",
	}

	fmt.Println(stus[1002].name)

	var stu student25
	fmt.Printf("%p\n", &stu.name)
	fmt.Printf("%p\n", &stu.age)
	fmt.Printf("%p\n", &stu.sex)
	fmt.Printf("%p\n", &stu.score)
	fmt.Printf("%p\n", &stu.addr)
}
