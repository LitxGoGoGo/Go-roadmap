package main

import "fmt"

func main1901() {
	//根据学生的成绩给出对应的评分
	//100 90 A; 80 B; 70 C;60 D;小于60 E
	var score int
	fmt.Scan(&score)

	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}

func main1902() {
	//根据学生的成绩给出对应的评分
	//100 90 A; 80 B; 70 C;60 D;小于60 E
	var score int
	fmt.Scan(&score)

	switch score / 10 {
	case 10:
		fallthrough //当case执行的代码过程中遇到了fallthrough 向下执行下一个case的内容
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
