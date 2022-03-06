package main

import (
	"fmt"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	//先创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//输出查看原始数组
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//遍历chessMap,如果有元素不为零那么就创建一个结构体,将其放入到对应的切片中

	var sparseArr []ValNode

	//标准的稀疏数组应该含有一个表示记录原始的二维数组的规模
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	//对chessMap进行遍历,找到对应的值
	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				//创建值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
		fmt.Println()
	}

	//输出稀疏数组
	for i, valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将稀疏数组存盘
	/*
		1.打开存盘文件
		2.恢复数组
	*/

	//恢复原始数组的操作
	var chessMap2 [11][11]int

	//遍历sparseArr [遍历文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.row] = valNode.val
		}
	}

	fmt.Println("恢复后的原始数据")
	for _, v1 := range chessMap2 {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
