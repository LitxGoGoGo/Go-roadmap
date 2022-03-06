package main

import (
	"fmt"
	"sync"
)

//定义双向链表的节点结构体类型

type DoubleNode struct {
	//数据域
	Data interface{}
	//指针域,两个指针,一个指向前一个节点,另一个指向后一个节点
	PreNode  *DoubleNode
	NextNode *DoubleNode
}

//定义双向链表数据结构的结构体

type DoubleLinkList struct {
	//链表头
	Head *DoubleNode
	//链表尾
	Tail *DoubleNode
	//链表长度
	Size int
	//读写锁
	mutex *sync.RWMutex
}

//创建链表
func (link *DoubleLinkList) CreateDoubleLinkList(Data ...interface{}) {
	//双向链表的初始化
	link.Size = len(Data)
	link.Head = new(DoubleNode)
	link.Tail = new(DoubleNode)
	link.mutex = new(sync.RWMutex)
	//读写锁的操作
	link.mutex.Lock()
	//解锁操作,defer延时调用
	defer link.mutex.Unlock()
	//设置头尾的值 用于最后赋值
	head := link.Head
	tail := link.Tail
	for i := 0; i < link.Size; i++ {
		//创建新节点并且添加数据
		//循环一个创建一个
		newlink := new(DoubleNode)
		newlink.Data = Data[i]
		//头节点的下一个节点
		link.Head.NextNode = newlink
		//将新节点的pre节点指向head节点
		newlink.PreNode = link.Head
		//激昂头节点替换成newlink
		link.Head = newlink
		tail = newlink
	}
	//整理双向链表头尾
	link.Head = head
	link.Tail = tail
}

//正序打印链表
func (link *DoubleLinkList) PrintDoubleLinkListByAse() {
	if link == nil {
		return
	}
	if link.Size == 0 {
		return
	}
	//读可以一起读,读取锁
	link.mutex.RLock()
	defer link.mutex.RUnlock()
	//头部节点不存储信息
	node := link.Head.NextNode
	for node != nil {
		fmt.Println(node.Data)
		node = node.NextNode
	}
}

//反序打印链表
func (link *DoubleLinkList) PrintDoubleLinkListByDesc() {
	if link == nil {
		return
	}
	if link.Size == 0 {
		return
	}
	//读可以一起读
	link.mutex.RLock()
	defer link.mutex.RUnlock()
	node := link.Tail
	for node.PreNode != nil {
		fmt.Println(node.Data)
		node = node.PreNode
	}
}

//指定位置插入对应数据
func (link *DoubleLinkList) InsertDataByIndex(index int, Date interface{}) {
	if link.Size == 0 {
		return
	}
	if index > link.Size || index < 0 {
		return
	}
	//写入锁打开
	link.mutex.Lock()
	defer link.mutex.Unlock()

	//记录指针信息
	//node 表示插入位置的原始节点
	node := link.Head
	for i := 0; i <= index; i++ {
		node = node.NextNode
	}
	//找到位置后创建新节点
	newnode := new(DoubleNode)
	newnode.Data = Date
	newnode.PreNode = node.PreNode
	newnode.NextNode = node
	node.PreNode.NextNode = newnode
	node.PreNode = newnode
}

//基于位置的查询
func (link *DoubleLinkList) SearchDataByIndex(index int) (result interface{}) {
	if index <= 0 {
		return nil
	}
	if link.Size < index {
		return nil
	}
	//添加读取锁
	link.mutex.RLock()
	defer link.mutex.RUnlock()
	//记录节点信息
	node := link.Head
	for i := 0; i < index; i++ {
		node = node.NextNode
	}
	return node.Data
}

//删除数据
func (link DoubleLinkList) DeleteDataByIndex(index int) {

}

//销毁链表
func (link DoubleLinkList) DestroyDoubleLinkList() {

}
