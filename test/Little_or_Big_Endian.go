package main

import (
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool { //小端模式:数据的高字节保存在内存的高地址中，而数据的低字节保存在内存的低地址中
	n := 0x1234
	ptr := (*[2]byte)(unsafe.Pointer(&n))
	fmt.Printf("%x\n", *ptr)
	if ptr[0] == 0x34 {
		return true
	}
	return false
}

func main() {
	if IsLittleEndian() == true {
		fmt.Println("小端")
	} else {
		fmt.Println("大端")
	}
}
