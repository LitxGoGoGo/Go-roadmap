package main

import (
	"errors"
	"fmt"
)

/*
一般分为三种异常
编辑时异常
编译时异常
运行时异常
*/

func result12(x int, y int) (value int, err error) {
	if y == 0 {
		err = errors.New("integer divide by zero")
		return
	}
	value = x / y
	return
}

func main1201() {
	a := 10
	b := 0
	value, err := result12(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}
