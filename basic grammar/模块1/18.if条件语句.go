package main

import "fmt"

func main1801() {
	var score int
	fmt.Scan((&score))

	if score > 700 {
		fmt.Println("我要上清华")
	} else {
		fmt.Println("我要上炕")
	}
}

func main1802() {
	var score int
	fmt.Scan(&score)

	//从上到下依次执行
	if score > 700 {
		fmt.Println("能上清华")
	} else if score > 690 {
		fmt.Println("能上北大")
	} else if score > 550 {
		fmt.Println("能上西大")
	} else {
		fmt.Println("能上炕")
	}

}
