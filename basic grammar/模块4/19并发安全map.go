package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	//将数据存储在sync map中
	m.Store(1001, "Litx")
	m.Store(1002, "张三")
	m.Store(1003, "李四")
	m.Store(1004, "王五")

	//var value, ok = m.Load(1111)
	//if ok {
	//	//类型断言
	//	val, ok := value.(string)
	//	if ok {
	//		fmt.Println(val)
	//	}
	//} else {
	//	fmt.Println("不存在对应值")
	//}

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		//为true后继续迭代遍历
		//为false则随机取出map中的值
		return false
	})
}
