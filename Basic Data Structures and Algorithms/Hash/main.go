package main

import (
	"fmt"
)

//定义emp
type EmpNode struct {
	id   int
	Name string
	Next *EmpNode
}

//定义哈希链表
//这里的链表不带表头,第一个节点直接存放雇员
type EmpLink struct {
	Head *EmpNode
}

//方法待定
//添加员工的方法
func (emplink *EmpLink) Insert(emp *EmpNode) {
	//辅助指针
	cur := emplink.Head
	var pre *EmpNode = nil
	//如果当前emplink为空链表那么直接把新进来的作为头节点
	if emplink.Head == nil {
		emplink.Head = emp
		return
	}
	//如果不是一个空链表
	for { //这个for循环是为了将指针定位到合适的位置
		if cur != nil {
			//比较
			if cur.id > emp.id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	emp.Next = cur
	pre.Next = emp

	//pre.Next = emp
	//emp.Next = cur
}

//显示节点的方法
func (emplink *EmpLink) ShowLinkNode(no int) {
	if emplink.Head == nil {
		fmt.Printf("当前链表%d为空\n", no)
		return
	}
	//不为空则遍历当前链所有节点进行显示
	temp := emplink.Head
	for {
		if temp != nil {
			fmt.Printf("当前链%d,雇员id为%d,名字%s->\n", no, temp.id, temp.Name)
			temp = temp.Next
		} else {
			break
		}
	}
	fmt.Println() //换行处理
}

//根据id查找链下对应的节点,没有返回空
func (emplink *EmpLink) FindById(id int) *EmpNode {
	cur := emplink.Head
	for {
		if cur != nil && cur.id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

//根据id找到链下对应的节点并进行删除
func (emplink EmpLink) DeleteById(id int) {
	//创建一个临时节点和
	constanstHead := emplink.Head
	var temp *EmpNode
	cur := emplink.Head
	//创建一个状态标识为flag,flag转为true则代表找到节点
	flag := false
	for {
		if cur == nil {
			fmt.Println("找不到该雇员")
			break
		} else if cur.id == id {
			flag = true //找到后改变状态
			break
		}
		temp = cur
		cur = cur.Next
	}
	if flag == true { //状态改变则说明找到
		//开始进行删除
		//不可删除头节点
		if cur == constanstHead {
			fmt.Println("头节点不能删寄寄")
		} else {
			temp.Next = cur.Next
			fmt.Println("删除成功")
		}
	}
}
func (e *EmpNode) ShowMe() {
	fmt.Printf("链表%d,找到该雇员%d\n,雇员名字%s\n", e.id%7, e.id, e.Name)
}

//定义hashtable,含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给hashtable添加雇员方法
func (has *HashTable) Insert(emp *EmpNode) {
	//使用散列函数确定雇员添加到哪一个链表
	linkNo := has.HashFun(emp.id)
	//使用对应的链表进行添加
	has.LinkArr[linkNo].Insert(emp)
}

//编写方法显示hashtable的所有雇员
func (has *HashTable) ShowAll() {
	for i := 0; i < len(has.LinkArr); i++ {
		has.LinkArr[i].ShowLinkNode(i)
	}
}

//编写一个用于散列的函数
func (has *HashTable) HashFun(id int) int {
	return id % 7 //对应链表的下表
}

//编写一个查找方法
func (has *HashTable) Find(id int) *EmpNode {
	//使用散列函数来查询雇员应该在哪一个链表
	linkNo := has.HashFun(id)
	return has.LinkArr[linkNo].FindById(id)
}

//编写删除方法
func (has *HashTable) DeleteById(id int) {
	linkNo := has.HashFun(id)
	has.LinkArr[linkNo].DeleteById(id)
}
func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("========雇员系统菜单========")
		fmt.Println("input 添加雇员")
		fmt.Println("show  显示雇员")
		fmt.Println("find  查找雇员")
		fmt.Println("del   删除雇员")
		fmt.Println("exit  退出系统")
		fmt.Println("请输入你的命令")
		fmt.Scan(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id")
			fmt.Scan(&id)
			fmt.Println("请输入雇员名字")
			fmt.Scan(&name)
			//构建雇员节点
			emp := &EmpNode{
				id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入你要查询的员工id")
			fmt.Scan(&id)
			emp := hashtable.Find(id)
			if emp == nil {
				fmt.Printf("id=%d的雇员信息不存在\n", id)
			} else {
				//编写方法显示雇员信息
				emp.ShowMe()
			}
		case "del":
			fmt.Println("请输入你要删除的雇员id")
			fmt.Scan(&id)
			hashtable.DeleteById(id)
		case "exit":
		default:
			fmt.Println("输入错误")
		}
	}
}
