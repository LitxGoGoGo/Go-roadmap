package main

import "fmt"

func main2001() {
	//命名规则  var 字典名 map[数据类型]值数据类型
	/*	var studentlist map[int]string = make(map[int]string)
		studentlist[1] = "zs"
		studentlist[2] = "ls"
		studentlist[3] = "ww"
		studentlist[4] = "zl"
		fmt.Println(studentlist)
	*/
	//	var studentlist map[int]string //这声明的map指向nil无法进行初始化操作
	/*		var studentlist = make(map[int]string)
			studentlist[1] = "zs"
			studentlist[2] = "ls"
			studentlist[3] = "ww"
			studentlist[4] = "zl"
			fmt.Println(studentlist)*/

	//使用自动推导类型进行创建并且赋予初始值
	studentlist := map[int]string{1: "zs", 2: "ls", 3: "ww", 4: "zl"}
	fmt.Println(studentlist)

	//可以使用rang对map进行遍历,输出的顺序是不固定的,随机的顺序
	for key, value := range studentlist {
		fmt.Println(key, value)
	}

}

func main2002() {
	//判断map中的值是否存在
	m := map[int]string{1: "刘能", 2: "赵四", 3: "谢广坤", 7: "王老七", 11: "王大拿"}

	value, ok := m[1]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("map中的key不存在")
	}
}
