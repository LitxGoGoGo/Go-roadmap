package main

import "fmt"

func add09(x, y int) int {
	return 0
}

func main0901() {
	//接口可以存储任意类型的数据,写hello就打印hello
	var i interface{} = add09
	var i09 interface{} = add09(1, 2)
	fmt.Printf("%T\n", i)   //函数类型 func(int ,int) int
	fmt.Printf("%T\n", i09) //int
	fmt.Println(i)          //函数的地址
}
