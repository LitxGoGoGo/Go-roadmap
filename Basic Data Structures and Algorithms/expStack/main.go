package main

import (
	"errors"
	"fmt"
	"strconv"
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

//判断一个字符是不是一个运算符
func (sta *Stack) IsOpe(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算方法
func (sta *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

//编写方法来展示优先级
func (sta *Stack) Priority(oper int) (res int) {
	res = 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

//
func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	////符号栈
	opeStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "3+2*6-2-87"
	//定义一个index,帮助扫描
	index := 0
	//为了配合运算,我们需要定义变量
	num1 := 0
	num2 := 0
	oper := 0
	keepNum := ""
	for {
		ch := exp[index : index+1]
		fmt.Println(ch)
		temp := int([]byte(ch)[0]) //转换成ASCII码值
		if opeStack.IsOpe(temp) {  //判断是不是符号
			//此时判断栈是否为空
			if opeStack.Top == -1 { //这是一个空栈
				opeStack.Push(temp)
			} else {
				//开始下一层逻辑,判断优先级
				if opeStack.Priority(opeStack.arr[opeStack.Top]) >= opeStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = opeStack.Pop()
					result := opeStack.Cal(num1, num2, oper)
					//得到的result重新入数栈
					numStack.Push(result)
					//把当前的符号压入符号栈
					opeStack.Push(temp)
				} else {
					opeStack.Push(temp)
				}

			}
		} else { //判断为数则压入数栈
			//设定一个变量,让他向后检查下一位是不是符号位,如果不是的化就合并
			keepNum += ch //先拼起来,关于下一个是否继续拼还得看下面的逻辑
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if opeStack.IsOpe(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		//继续扫描
		//先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++ //保证继续扫描
	}
	//扫描玩表达式后,依次从符号栈中取出符号,然后从数栈中取出两个数
	//运算后的结果,入数栈,直到符号栈为空
	for {
		if opeStack.Top == -1 {
			break //退出条件
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = opeStack.Pop()
		result := opeStack.Cal(num1, num2, oper)
		//将结果入数栈
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s=%d", exp, res)
}
