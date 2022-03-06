package main

import (
	"fmt"
	"math/rand"
	"time"
)

//编写选择排序函数,go中为值传递所以要通过指针来改变原数组的值
func SelectSort(arr *[160000]int) {
	//假设arr[0]就是最大值
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		//遍历数组,有比max大的就进行替换
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		//fmt.Printf("第%d次 %v\n", j+1, *arr)
	}
}

func main() {
	//定义一个数组
	//arr := [5]int{10, 34, 19, 100, 80}
	//从大到小排序
	var arr [160000]int
	for i := 0; i < 160000; i++ {
		arr[i] = rand.Intn(90000)
	}
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("耗时%d秒", end-start)

}
