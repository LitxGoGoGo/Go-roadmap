package main

import "fmt"

func test19(slice []int) {
	slice[1] = 111
	fmt.Println("被调函数", slice)
}

func main1901() {
	slice := []int{1, 2, 3, 4, 5}
	test19(slice)
	fmt.Println("主函数原始切片的值", slice)
}

//对切片元素进行排序
func main1902() {
	slice := []int{2, 4, 6, 8, 9, 1, 7, 5, 3, 10}
	BubbleSort1(slice)
	fmt.Println(slice)
}

//使用冒泡排序法
func BubbleSort1(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
