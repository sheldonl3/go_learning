package main

import (
	"fmt"
	"sync"
)

var num int
var once sync.Once //方法内执行代码只执行一次

func add_1() {
	num += 1
	fmt.Println("add 1")
	return
}
func add_2() {
	num += 1
	fmt.Println("add 2")
	return
}
func main() {
	for i := 0; i < 10; i++ {
		once.Do(add_1)
	}
	once.Do(add_2)
	fmt.Println(num)
	return
}
