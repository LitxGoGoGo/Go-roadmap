package main

import "fmt"

// Student24 定义结构体
type Student24 struct {
	id    int
	name  string
	age   uint8
	sex   rune
	score float64
	addr  string
}

// BubbleSort24 用冒泡排序来将结构体切片按照成绩进行排序
func BubbleSort24(stu []Student24) {
	for i := 0; i < len(stu); i++ {
		for j := 0; j < len(stu)-1-i; j++ {
			if stu[j].score < stu[j+1].score {
				stu[j], stu[j+1] = stu[j+1], stu[j]
			}
		}
	}
}

func main2401() {
	//定义结构体切片
	var stus = []Student24{{
		id:    1001,
		name:  "刘能",
		age:   52,
		sex:   '男',
		score: 88,
		addr:  "象牙山27号",
	}, {
		id:    1052,
		name:  "谢广坤",
		age:   51,
		sex:   '男',
		score: 59.5,
		addr:  "象牙山78号",
	}, {
		id:    1092,
		name:  "王老七",
		age:   55,
		sex:   '男',
		score: 99.5,
		addr:  "象牙山1号",
	}}

	//通过结构体切片来修改结构体成员的信息
	//stus[1].score = 10

	BubbleSort24(stus)

	for _, value := range stus {
		fmt.Println(value)
	}
}
