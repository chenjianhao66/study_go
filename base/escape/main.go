package main

import "fmt"

// 1. 指针对象逃逸

type Student struct {
	Name string
	Age  int
}

func NewStudent(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

// 2. 栈空间不足逃逸

func slice() {
	s := make([]int, 1000, 10000)

	for i, _ := range s {
		s[i] = i
	}
}

func main() {
	// 1. 指针对象逃逸
	NewStudent("jack", 25)

	// 2. 栈空间不足逃逸

	slice()

	// 3. 动态类型逃逸
	s := "escape"
	fmt.Println(s)
}
