package main

import "fmt"

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func main1301() {
	s := make(Set)
	s.Add("Litx")
	s.Add("李四")
	fmt.Println(s.Has("Litx"))
	fmt.Println(s.Has("李四"))
	fmt.Println(s.Has("张三"))
}

type Student struct {
	name string
	age  int
}

func main1302() {
	var s1 = new(Student)
	s1.age = 22
	s1.name = "tcy"
	var s2 = Student{age: 21, name: "tcy"}
	var s3 = &s2
	var s4 = &s3
	fmt.Println(s1, s2) //&{tcy 22} {tcy 21}
	fmt.Printf("%T\n%T\n", s1, s2)
	fmt.Println(&s1, s4) //可以看出计算机操作系统为s1和s2开辟了不同的内存空间
}
