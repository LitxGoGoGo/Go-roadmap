package main

import "fmt"

func main() {
	value := sum(1, 2)
	fmt.Println(value)
	fmt.Printf("%T", value)

	value = sum(1, 2, 3, 4, 5)
	fmt.Println(value)
	fmt.Printf("%T", value)
}
func sum(args ...int) int {
	value := 0
	for i := 0; i < len(args); i++ {
		value += i
	}
	return value
}
