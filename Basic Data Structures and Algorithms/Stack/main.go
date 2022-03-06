package main

import (
	"errors"
	"fmt"
)

//定义栈的结构体
type Stack struct {
	MaxTop int
	Top    int //因为栈底是固定的,因此使用Top就行
	arr    [5]int
}

//操作栈的函数,入栈
func (sta *Stack) Push(val int) (err error) {
	//先判断栈是否已经满了
	if sta.Top == sta.MaxTop-1 {
		fmt.Println("栈已满")
		return errors.New("stack full")
	}
	//先让Top+1 ,然后放入数据
	sta.Top++
	sta.arr[sta.Top] = val
	return
}

//出栈函数
func (sta *Stack) Pop() (val int, err error) {
	//判断是否为空
	if sta.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("stack empty")
	}
	//先取值
	val = sta.arr[sta.Top]
	fmt.Printf("值%d出栈\n", val)
	sta.Top--
	return val, nil
}

//遍历栈,需要从栈顶开始遍历
func (sta *Stack) List() {
	//先判断栈是否为空
	if sta.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	//curTop := sta.Top
	for i := sta.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, sta.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5) //栈顶
	stack.Push(6)

	//显示
	stack.List()

	val, _ := stack.Pop()
	fmt.Println(val)
	stack.List()
}
