package main

import "fmt"

//定义函数类型
type addfn func(a int, b int) int

//在同一个包下函数名唯一
func add(x int, y int) int {
	sum := x + y
	return sum
}

func main0701() {

	a := 10
	b := 20
	result := add(a, b)
	fmt.Println(result)

	//定义函数类型后使用
	x := 20
	y := 30
	var af addfn = add
	value := af(x, y)
	fmt.Println(value)

	//函数类型作为参数进行传递
	p := 30
	q := 40
	res := jiafa(p, q, add)
	fmt.Println(res)
}

//将函数类型作为函数的参数进行传递
func jiafa(m int, n int, mimi addfn) int {
	return mimi(m, n)
}
