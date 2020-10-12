package main

import (
	"algorithm/bstree"
	"algorithm/bstree/travel"
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
	root := bstree.Build_tree("123#45")
	res := make([]int, 0)
	travel.Postorder(&root, &res)
	fmt.Println(res)
}
