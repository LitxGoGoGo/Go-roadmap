package main

import "fmt"

func main0601() {
	var i int
	var p *int
	var q **int
	p = &i
	q = &p
	w := &q
	i = 1
	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &p)
	fmt.Printf("%p\n", q)
	fmt.Printf("%p\n", &q)
	fmt.Printf("%p\n", w)
	fmt.Printf("%p\n", &w)

	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", *p)
	*p = 9999
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", *p)
	**q = 88
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", *p)

	fmt.Printf("%p\n", &i)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &p)
	fmt.Printf("%p\n", q)
	fmt.Printf("%p\n", &q)
	fmt.Printf("%p\n", w)
	fmt.Printf("%p\n", &w)

}
