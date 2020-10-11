package main

import (
	"fmt"
)

func ScanLine() string { //输入一行字符串，包括空格
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)
		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

func main() {
	fmt.Println(uniquePaths(7, 3))
}
