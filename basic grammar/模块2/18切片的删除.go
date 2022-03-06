package main

import "fmt"

func main1801() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(slice), cap(slice)) //len 10 cap 10
	//容量只会变大不会变小
	//切片的删除 结果：[1 2 8 9 10]
	slice = append(slice[:2], slice[7:]...)
	fmt.Println(len(slice), cap(slice))

	fmt.Println(slice)

}
