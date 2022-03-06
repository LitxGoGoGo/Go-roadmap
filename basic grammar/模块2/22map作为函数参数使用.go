package main

import "fmt"

//创建修改函数
func modify(m map[int]string) {
	//修改操作
	m[2] = "222"
	//添加操作
	m[5] = "555"
}

func main2201() {
	m := map[int]string{1: "刘能", 2: "赵四", 3: "谢广坤", 10: "王老七"}
	fmt.Println(m)
	//在被调函数中修改会影响主调函数中map的值
	modify(m)

	fmt.Println(m)
}
