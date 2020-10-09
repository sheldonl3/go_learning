package main

import (
	"fmt"
	"time"
)

func f(n int) (r int) {
	fmt.Println("r", r)
	defer func() {
		fmt.Println("r", r) //此时变成return的值
		if recover() != nil {
			r += n
			fmt.Println("hello")
		}
	}()
	var f func()
	defer f() //panic
	f = func() {
		r += 2
		fmt.Println("world")
	}
	return n + 1
}

func test_panic() string {
	s := "hgj"
	defer func() {
		recover()
	}()
	panic("sdd")
	return s
}

func test2_panic() (res string) {
	res = test_panic()
	return res
}

func test_goroute() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Duration(2) * time.Second)
}

func istmp() { //range的list是临时变量,slice就是引用变量
	var lis = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range lis {
		if k == 0 {
			fmt.Println("me", v)
			lis[6] = 1000
			fmt.Println(lis)
		} else {
			fmt.Println(v)
		}

	}
}

func main() {
	//fmt.Println(f(3))
	//test_goroute()
	istmp()
}
