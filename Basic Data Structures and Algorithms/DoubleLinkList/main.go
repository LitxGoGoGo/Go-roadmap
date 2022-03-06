package main

import "fmt"

//定义双向链表节点结构体信息
type HeroNode struct {
	no       uint
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

//给链表插入一个节点
//编写第一种插入方法
func InsertHeroNode(headNode *HeroNode, newHeroNode *HeroNode) {
	//1,先找到该链表最后一个节点
	//2,创建一个辅助节点
	temp := headNode
	for {
		if temp.next == nil {
			break //跳出循环则说明找到了链尾的指针
		}
		temp = temp.next //没到最后就让temp一直加
	}

	//3,让新节点加到最后一个节点的next中
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//编写第二种插入方法,按照编号从小到大排序
func InsertHeroNode2(headNode *HeroNode, newHeroNode *HeroNode) {
	//1,找到适当的节点
	//2,创建一个辅助节点
	temp := headNode
	flag := true
	//让新节点的no和temp节点的下一个节点的no进行比较
	for {
		if temp.next == nil { //说明到链表的最后了
			break
		} else if temp.next.no > newHeroNode.no {
			//说明这个新的node就应该插入temp的后边
			break
		} else if temp.next.no == newHeroNode.no {
			//说明已经有这个no了
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("该人物的no=", newHeroNode.no, "已经存在")
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

//双向链表删除操作,需要提供头节点和要删除的no
func Delet(headNode *HeroNode, id uint) {
	//创建一个临时节点temp
	temp := headNode
	flag := false
	//往下遍历找到对应id节点的上一个节点
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			//找到后改变flag的状态
			flag = true
			break
		}
		temp = temp.next
	}

	if flag == true {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的id不存在")
	}
}

//显示链表的所有信息
func ListHeroNode(headNode *HeroNode) {
	//1,创建辅助节点
	temp := headNode
	//先判断该链表是不是空链表
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}

	//2,遍历链表
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//	3,判断是不是链表尾部
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//从尾部向前展示链表信息
func ListHeroNode2(headNode *HeroNode) {
	//1,创建辅助节点
	temp := headNode

	//先判断该链表是不是空链表
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}

	//2,让temp找到表尾最后一个节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//3,遍历链表
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.no, temp.name, temp.nickname)
		//	3,判断是不是链表头部
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}
func main() {

	//1,先创建头节点,头节点为空没问题
	head := HeroNode{}

	//2,创建新节点HeroNode
	hero1 := HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	//加入
	InsertHeroNode2(&head, &hero3)
	InsertHeroNode2(&head, &hero1)
	InsertHeroNode2(&head, &hero2)
	//显示
	ListHeroNode(&head)
	fmt.Println()
	ListHeroNode2(&head)
	fmt.Println()
	Delet(&head, 1)
	ListHeroNode(&head)
}
