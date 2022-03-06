package main

import "fmt"

//计算N个数字的和
func add10(value ...int) int {
	fmt.Println(value)
	fmt.Printf("%T\n", value) //[]int 切片

	var sum int //默认值为0
	//遍历一个数据集合
	for _, result := range value {
		sum += result
	}
	fmt.Println(sum)
	return 0
}

func main1001() {
	add10(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
}
