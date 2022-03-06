package main

import (
	"fmt"
	"reflect"
)

type LinkList struct {
	//指针域
	NextNode *LinkList
	//数据域
	Data interface{}
}

//链表功能函数
//创建链表
func (link *LinkList) CreateLinkList(Data ...interface{}) {
	head := link
	for i := 0; i < len(Data); i++ {
		newlink := new(LinkList)
		//为新的节点添加数据
		//为节点添加数据
		newlink.Data = Data[i]
		//将节点的指针域指向新的节点
		link.NextNode = newlink
		link = newlink
	}
	link = head
}

//打印链表
func (link *LinkList) PrintLinkList() {
	if link == nil {
		return
	}
	if link.Data != nil {
		fmt.Println(link.Data)
	}
	//通过递归查找下一个节点的数据
	link.NextNode.PrintLinkList()
}

//获取链表节点个数
func (link *LinkList) GetLinkListNode() (count int) {
	if link == nil {
		return 0
	}
	for link != nil {
		count++
		link = link.NextNode
	}
	//链表头没有传输数据
	return count - 1
}

//指定位置插入数据
func (link *LinkList) InsertDataByIndex(index int, Data interface{}) {
	if link == nil {
		return
	}
	//如果插入位置超出链表长度直接返回
	if link.GetLinkListNode() < index {
		return
	}
	//插入位置合适则通过遍历找到插入的位置
	for i := 0; i < index; i++ {
		if link.NextNode != nil {
			link = link.NextNode
		}
	}
	//创建新节点
	newlink := new(LinkList)
	newlink.Data = Data //将数据存入新节点数据域
	//将新节点的下一个指针域指向之前节点的指针
	newlink.NextNode = link.NextNode
	link.NextNode = newlink

}

//在头部插入数据
func (link *LinkList) InsertDataByHead(Data interface{}) {
	link.InsertDataByIndex(0, Data)
}

//在尾部插入数据
func (link *LinkList) InsertDataByTail(Data interface{}) {
	link.InsertDataByIndex(link.GetLinkListNode(), Data)
}

//查找数据
func (link *LinkList) SearchData(Data interface{}) (result interface{}) {
	if link == nil {
		return
	}
	//通过反射判断两个接口是否一致,同时判断内容是否一致
	if reflect.TypeOf(link.Data) == reflect.TypeOf(Data) && link.Data == Data {
		return link.Data
	}
	return link.NextNode.SearchData(Data)
}

//删除指定位置数据
func (link *LinkList) DeleteDataByIndex(index int) {
	if link == nil {
		return
	}
	//判断index的取值范围
	if link.GetLinkListNode() < index {
		return
	}
	//记录下被删位置前一个节点
	prev := link
	for i := 0; i < index; i++ {
		if link.NextNode != nil {
			prev = link
			link = link.NextNode
		}
	}
	prev.NextNode = link.NextNode
	//将删除的节点数据和指针域都置为空
	link.Data = nil
	link.NextNode = nil
}

//删除指定值的数据(在一个链表中删除第一次出现的数据)
func (link *LinkList) DeleteDataByValue(Data interface{}) {
	prev := link
	for link != nil {
		if reflect.TypeOf(link.Data) == reflect.TypeOf(Data) && link.Data == Data {
			prev.NextNode = link.NextNode

			link.Data = nil
			link.NextNode = nil
		}
		prev = link
		link = link.NextNode
	}
}

//销毁链表
func (link *LinkList) DestroyLinkList() {
	if link == nil {
		return
	}
	link.NextNode.DestroyLinkList()
	link.Data = nil
	link.NextNode = nil
}
