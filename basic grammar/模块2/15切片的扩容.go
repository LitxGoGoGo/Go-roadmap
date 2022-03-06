package main

import "fmt"

func main1501() {
	//创建切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//使用append对切片进行扩容
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//在切片进行扩容时，如果小于1024 每次扩容为上一次的2倍 如果大于等于1024 每次扩容为上一次的1/4

}
