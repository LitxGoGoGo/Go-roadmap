package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr *[160000]int) {
	for j := 1; j < len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j - 1

		//从大到小排序
		for insertIndex >= 0 && insertVal > arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex] //后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}
		//fmt.Printf("第%d次插入后%v\n", j, *arr)
	}
}

func main() {
	//arr := [5]int{23, 0, 12, 56, 34}
	var arr [160000]int
	for i := 0; i < 160000; i++ {
		arr[i] = rand.Intn(90000)
	}
	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("耗时%d秒", end-start)
	fmt.Println("main主函数")

}
