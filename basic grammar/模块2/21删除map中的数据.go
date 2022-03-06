package main

import "fmt"

func main2101() {
	m := map[int]string{3: "李四", 5: "小淘气", 12: "张三", 1: "王五"}
	fmt.Println(m)

	//删除操作使用方法delete
	//数据存在则删除,不存在则跳过
	delete(m, 12)
	fmt.Println(m)

}
