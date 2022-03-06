package main

import "fmt"

func main1601() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//切片[起始位置,结束位置] 左闭右开 包含起始抛弃末尾
	s := slice[0:3] //截取内容[1,2,3]
	fmt.Println(s)
	//修改切片s 会影响切片slice
	s[1] = 111
	fmt.Println(s)
	fmt.Println(slice)
}

func main1602() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//使用[:]获取原有切片的所有值
	s1 := slice1[:]
	fmt.Printf("变量slice地址: %p\n", &slice1)
	fmt.Printf("变量s地址: %p\n", &s1)
}
