package main

import "fmt"

func add1101(x int, y int) int {
	sum := add1102(x, y)
	return sum
}
func add1102(x int, y int) int {
	return x + y
}
func main1101() {
	a := 20
	b := 30
	sum := add1101(a, b)
	fmt.Println(sum)
}

//全局变量,首字母大写作用域为整个项目,首字母小写作用域为整个包
var sum int = 1

//递归函数计算阶乘
func add1103(value int) {
	if value == 1 {
		return
	} else {
		sum *= value
		add1103(value - 1)
	}
}

func main1102() {
	add1103(8)
	fmt.Println(sum)
}
