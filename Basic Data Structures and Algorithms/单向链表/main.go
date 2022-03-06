package main

import "fmt"

type Student struct {
	id    int
	name  string
	sex   string
	score float64
	addr  string
}

func main() {
	var linklist *LinkList = new(LinkList)

	//创建链表
	linklist.CreateLinkList(Student{
		id:    1001,
		name:  "刘能",
		sex:   "男",
		score: 90,
		addr:  "象牙村70号",
	}, Student{
		id:    1005,
		name:  "谢广坤",
		sex:   "男",
		score: 59,
		addr:  "象牙村10号",
	}, Student{
		id:    1001,
		name:  "刘能",
		sex:   "男",
		score: 90,
		addr:  "象牙村70号",
	}, Student{
		id:    1020,
		name:  "赵四",
		sex:   "男",
		score: 88,
		addr:  "象牙村35号",
	}, Student{
		id:    1001,
		name:  "刘能",
		sex:   "男",
		score: 90,
		addr:  "象牙村70号",
	})
	//打印链表
	linklist.PrintLinkList()
	//计数
	count := linklist.GetLinkListNode()
	fmt.Println("总数为", count)
	//指定位置插入数据
	linklist.InsertDataByIndex(3, Student{
		id:    1030,
		name:  "王老七",
		sex:   "男",
		score: 99,
		addr:  "象牙山80",
	})
	linklist.InsertDataByIndex(5, Student{
		id:    1030,
		name:  "王老七",
		sex:   "男",
		score: 99,
		addr:  "象牙山80",
	})
	linklist.PrintLinkList()

	//查找赵四的信息
	//stuZS := Student{
	//	id:    1020,
	//	name:  "赵四",
	//	sex:   "男",
	//	score: 88,
	//	addr:  "象牙村35号",
	//}
	//result := linklist.SearchData(stuZS)
	//if result != nil {
	//	fmt.Println("xxxxxxxxxxxxxxxxxx")
	//	fmt.Println(result)
	//} else {
	//	fmt.Println("未找到数据")
	//}

	/*删除王老七的数据*/
	stuWLQ := Student{
		id:    1030,
		name:  "王老七",
		sex:   "男",
		score: 99,
		addr:  "象牙山80",
	}

	linklist.DeleteDataByValue(stuWLQ)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXX")
	linklist.PrintLinkList()

	/*销毁链表*/
	linklist.DestroyLinkList()
	linklist.PrintLinkList()
}
