package main

import (
	"fmt"
)

func main1401() {
	//切片的创建方式
	//var slice []int //空切片
	//slice[0] = 1 //这会报错,切片元素个数为零 不能存储任何数据
	//fmt.Println(slice)

	//使用make创建切片,make([]数据类型,长度,容量)
	slice := make([]int, 5, 10) //未初始化的会按照类型默认的初始值进行初始化
	slice[0] = 20
	slice[1] = 10
	fmt.Println(slice)

	//使用自动推导类型进行切片的创建
	slice1 := []int{1: 11, 2: 22, 4: 44} //可以进行初始化
	//使用自带的len查询长度,cap查询容量
	fmt.Println(len(slice1)) //长度为5
	fmt.Println(cap(slice1)) //容量为5
	fmt.Println(slice1)
}

func main1402() {
	//遍历切片
	slice3 := make([]int, 5, 10)
	slice3 = []int{1, 2, 3, 4, 5}

	//使用range遍历
	for index, value := range slice3 {
		fmt.Println(index, value)
	}
}
