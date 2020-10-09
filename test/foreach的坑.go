package main

import (
	"fmt"
	"unsafe"
)

type student struct {
	Name string
	Age  int
}

func wrong_way(m map[string]*student) {
	stus := []student{
		{Name: "zhou", Age: 23},
		{Name: "ewrf", Age: 42},
		{Name: "yur", Age: 12},
	}
	for _, stu := range stus { //for each 迭代切片，不能通过迭代器的使用每一项内容的地址: 迭代器的指针，最后都指向了最后一个值
		fmt.Printf("%v\n",unsafe.Pointer(&stu) )
		m[stu.Name] = &stu
	}
	return
}

func right_way(m map[string]*student) {
	stus := []student{
		{Name: "zhou", Age: 23},
		{Name: "ewrf", Age: 42},
		{Name: "yur", Age: 12},
	}
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	return
}

func main() {
	m := make(map[string]*student)
	wrong_way(m)
	fmt.Println("wrong_way")
	for name, stu := range m {
		fmt.Println(name, stu)
	}

	right_way(m)
	fmt.Println("\nright_way")
	for name, stu := range m {
		fmt.Println(name, stu)
	}
}
