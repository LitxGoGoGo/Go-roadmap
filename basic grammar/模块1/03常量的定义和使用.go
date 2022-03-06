package main

import "fmt"

func main0301() {
	//常量的定义和使用
	const A int = 10
	//A = 100 //常量在运行过程中不能进行赋值操作
	fmt.Println(A)
}

func main0302() {
	const (
		A = 11
		B = 22
		C = 33
	)
	fmt.Println(A, B, C)
}

func main0303() {
	const (
		A, D = iota, iota //0 状态会改变,同一行的值为相同
		B, E = iota, iota //1 换行+1
		C, F = iota, iota //2
	)
	fmt.Println(A, B, C, D, E)

}
