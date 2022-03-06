package main

import "fmt"

type BoyNode struct {
	No   int
	next *BoyNode
}

//num表示数的数,最终函数返回环形链表指向第一个小孩的指针
func AddBoy(num int) *BoyNode {
	first := &BoyNode{}
	curBoy := &BoyNode{}
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	//循环构建这个循环链表
	for i := 1; i <= num; i++ {
		boy := &BoyNode{
			No: i,
		}
		//需要一个辅助指针
		//第一个小孩的位置比较特殊
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first //在只有一个小孩的时候自身构成循环
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first //构成环形链表
		}
	}
	return first
}

//显示单向环形链表
func Show(first *BoyNode) {
	if first.next == nil {
		fmt.Println("链表为空.没有小孩")
		return
	}
	//定义一个指针帮助遍历,到了这步肯定至少有一个小孩
	curBoy := first
	for {
		fmt.Printf("小孩的编号=%d->", curBoy.No)
		if curBoy.next == first {
			break
		}
		curBoy = curBoy.next
	}

}

//编写函数来解决问题
//1,函数PlayGame(first *Boy,startNo int,countNum int)
//使用一个算法让链表只剩下最后一个人
func PlayGame(first *BoyNode, startNo int, countNum int) {
	// 1,空的链表单独处理
	if first.next == nil {
		fmt.Println("空的链表,没有小孩")
		return
	}

	//判断的startNo要小于等于小孩的总数
	//2,定义一个指针辅助删除
	tail := first
	//3,让tail指向环形链表的最后一个小孩
	for {
		if tail.next == first {
			break
		} //说明到了最后一个小孩
		tail = tail.next
	}
	//4,让first移动到startNo[删除时以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	//5,开始数countNo,然后删除first指向的节点
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("小孩编号为%d的出列->\n", first.No)
		//删除first指向的节点
		first = first.next
		tail.next = first
		//判断如果tail == first 圈中只有一个小孩
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩编号是%d\n", first.No)

}

func main() {
	first := AddBoy(50)
	Show(first)
	fmt.Println()
	PlayGame(first, 19, 8)
}
