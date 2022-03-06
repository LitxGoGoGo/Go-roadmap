package main

import "fmt"

//地图函数,为了保证是同一个地图则使用*地址进行操作
func SetWay(myMap *[8][7]int, i int, j int) bool {
	//如果map[6][5] == 2那么就代表找到了
	if myMap[6][5] == 2 {
		return true
	} else {
		//进入这里则要继续找
		if myMap[i][j] == 0 {
			//这是能够进行探测的
			myMap[i][j] = 2
			//寻路策略下右上左
			if SetWay(myMap, i+1, j) { //下
				return true
			} else if SetWay(myMap, i, j+1) { //右
				return true
			} else if SetWay(myMap, i-1, j) { //上
				return true
			} else if SetWay(myMap, i, j-1) { //左
				return true
			} else { //上下左右都不同已经围住了
				myMap[i][j] = 3
				return false
			}
		} else {
			//说明这个点不能探测

			return false
		}
	}
}

func main() {
	//创建二维数组,模拟迷宫
	//元素的值为1代表墙(走不了)
	//元素的值为0则代表没有走过的点
	//元素的值变为2则代表可以通行
	//元素3表示曾经走过但是走不通
	var myMap [8][7]int

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	for i := 0; i < 7; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[1][2] = 1
	myMap[2][2] = 1
	myMap[3][1] = 1
	myMap[3][2] = 1
	//输出地图
	for _, val := range myMap {
		for _, i3 := range val {
			fmt.Printf("%d ", i3)
		}
		fmt.Println()
	}
	//使用测试
	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕")
	for _, val := range myMap {
		for _, i3 := range val {
			fmt.Printf("%d ", i3)
		}
		fmt.Println()
	}
}
