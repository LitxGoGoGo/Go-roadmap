package main

import "fmt"

func main0801() {
	a := 10
	b := 20
	//匿名函数
	sum := func(x int, y int) int {
		return x + y
	}(a, b) //括号表示匿名函数调用

	fmt.Println(sum)
}

func main0802() {
	a := 10
	b := 20
	//在当前函数中多次调用匿名函数
	add := func(x int, y int) int {
		return x + y
	} //定义匿名函数后赋值给了add

	sum := add(a, b) //调用了函数类型变量,该变量指向函数中的匿名函数
	fmt.Println(sum)
	fmt.Println(add(a, b))
	fmt.Printf("%T", add(a, b))
}
