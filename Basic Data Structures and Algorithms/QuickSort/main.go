package main

import "fmt"

//快速排序
func QuickSort(left int, right int, arr *[9]int) {
	l := left
	r := right
	pivot := arr[(left+right)/2] //中轴
	temp := 0

	//for循环的目标为将比pivot小的放左边,比pivot大的放右边
	for l < r {
		for arr[l] > pivot {
			l++
		}
		for arr[r] < pivot {
			r--
		}
		if l >= r {
			break
		}
		//交换
		temp = arr[l]
		arr[l] = arr[r]
		arr[r] = temp
	}
	//如果 l==r 再移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	arr := [9]int{-9, 78, 0, 23, -567, 70, 99, 100, 189}

	//调用排序
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main")
	fmt.Println(arr)
}
