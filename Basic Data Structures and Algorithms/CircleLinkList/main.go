package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

//添加猫猫
func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //初始化的时候应该指向自己后形成环形
		fmt.Println(newCatNode, "加入环形链表中")
		return //没有这个return没办法跳出
	}

	//定义临时变量,找到环形链表最后的节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入队列中
	temp.next = newCatNode
	newCatNode.next = head
	value1 := temp.next
	fmt.Println("新猫", value1.no, "已加入")

}

//输出环形链表
func ListCircleLinkList(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表空,不输出哼")
		return
	}
	for {
		fmt.Printf("小猫的信息是=[id=%d,name=%s]->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

//删除节点
func Delete(head *CatNode, id int) *CatNode {
	//temp指向头节点,helper指向环形链表的最后,让temp和要删除的id进行比较如果相同则使用helper来进行操作
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("链表空,不输删除")
		return head
	}
	//如果只有一个节点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil //被go的垃圾回收机制回收
		}
		fmt.Println("表头是否为空?", head == nil)
		return head
	}
	//将helper定位到链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有两个以上的节点,用for循环找到要删除的前一个节点
	flag := true
	for {
		if temp.next == head { //找到这里说明已经比较到最后一个了
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的是头节点
				head = head.next
			}
			//这表示找到了节点并且temp处于该节点
			helper.next = temp.next
			fmt.Printf("删除的猫猫=%d\n", id)
			flag = false

			break
		}
		temp = temp.next     //用来比较
		helper = helper.next //用来操作删除节点
		//如果flag为真那么表示还没有删除过继续下面的判断
		if flag {
			if temp.no == id {
				helper.next = temp.next
				fmt.Printf("删除猫猫=%d\n", id)
			} else {
				fmt.Println("没有此id的猫")
			}
		}
	}
	return head
}

func main() {
	//初始化环形链表头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLinkList(head)

	Delete(head, 3)
	ListCircleLinkList(head)

}
