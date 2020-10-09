package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

type Books struct {
	price  int16
	y      bool   //字节对齐，最后的结构体是最大属性内存大小的倍数;bool放到最后size就不一样
	author string //string又一个指向字符的指针和string大小组成,所以大小是总是16
	d      int
}

func main() {
	book := Books{12, true, "misaka", 45}
	fmt.Println(a, b, c, unsafe.Sizeof(book.author))
	fmt.Println("\nbook size ", unsafe.Sizeof(book))
	fmt.Println("book.price size ", unsafe.Sizeof(book.price))
	fmt.Println("book.author size", unsafe.Sizeof(book.author))
	fmt.Println("book.d size", unsafe.Sizeof(book.d))
	fmt.Println("book.y size", unsafe.Sizeof(book.y))
	fmt.Println(book)
	ptrstr := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&book)) + unsafe.Offsetof(book.author)))
	//一个unsafe.Pointer指针也可以被转化为uintptr类型，然后用以做必要的指针数值运算。
	//之后再转化成unsafepointer->再转化成具体类型的指针，就可以赋值操作
	*ptrstr = "kik"
	fmt.Println(book)
}
