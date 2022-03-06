package main

import "fmt"

func main1701() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := make([]int, 10, 10)
	//copy(目标切片,源切片)
	copy(dst, slice)
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", dst)
	dst[3] = 2333
	fmt.Println(slice)
	//拷贝后的切片修改并不会影响原来的切片
	fmt.Println(dst)

}

func main1702() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := make([]int, 9, 50)
	//如果源切片大于目标切片  会将部分切片进行拷贝
	//如果源切片小于目标切片  会将源切片所有内容拷贝
	copy(dst, slice)
	fmt.Println(dst)
}
