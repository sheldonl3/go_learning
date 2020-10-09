package main

import (
	"fmt"

)

//string []byte
//string在内存中的存储结构是长度固定的字节数组，也就是说是字符串是不可变的。
//当要修改字符串的时候，需要转换为[]byte，修改完成后再转换回来。
func string_to_byte(str string) []byte{
	by:=[]byte(str)
	fmt.Println(len(by))
	return by
}

func byte_to_string(lis []byte)string{
	res:=string(lis)
	return res
}


//中文
func string_to_rune(str string) []rune{
	by:=[]rune(str)
	fmt.Println(len(by))
	return by
}

func main() {
	fmt.Println(byte_to_string(string_to_byte("test按时")))
	fmt.Println(string_to_rune("test按时"))
}
