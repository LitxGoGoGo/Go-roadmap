package main

import "fmt"

func test12(x int, y int) {

}

type struct12 struct {
}

func mai1201() {
	var i interface{} = 10 //将整型赋值给interface{}
	var i2 interface{} = 3.14
	var i3 interface{} = "hello"
	var i4 interface{} = [3]int{1, 2, 3}
	var i5 interface{} = test12
	var i6 interface{} = struct12{}

	fmt.Println(i, i2, i3, i4, i5, i6)

	if value, ok := i2.(int); ok {
		fmt.Println(value)
		fmt.Printf("%T\n", value)
	} else {
		fmt.Println("不是int数据类型")
	}

	//使用switch进行流程控制
	/*	switch i2 {
		case i2.(int):
			fmt.Println("是int类型")
		case i2.(float64):
			fmt.Println("是浮点型")
		default:
			fmt.Println("未知类型")

		}*/ //有待验证
}
