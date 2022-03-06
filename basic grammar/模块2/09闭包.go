package main

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main0901() {
	fn := getSequence()
	//调用闭包的操作
	//value := fn()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())

	fn1 := getSequence()
	fmt.Println(fn1())
	fmt.Printf("%T\n", fn1)
	//fn1赋值后类型为func() int,函数类型
	fmt.Println(fn1())
	fmt.Println(fn1())
}
