package main

import "fmt"

func main() {
	var dlinklist *DoubleLinkList = new(DoubleLinkList)
	//创建双向链表
	dlinklist.CreateDoubleLinkList(1, 2, 3, 4, 5)
	//正序打印
	//dlinklist.PrintDoubleLinkListByAse()
	//反序打印
	//dlinklist.PrintDoubleLinkListByDesc()
	//fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	dlinklist.InsertDataByIndex(0, 6)
	dlinklist.PrintDoubleLinkListByAse()
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	value := dlinklist.SearchDataByIndex(4)
	if value != nil {
		fmt.Println(value)
	} else {
		fmt.Println("未找到数据")
	}
}
