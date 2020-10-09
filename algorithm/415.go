package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	fmt.Println(reflect.TypeOf(num1[0]))
	res := ""
	jinwei := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || jinwei != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0') //asc to int
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		wei := x + y + jinwei
		res = strconv.Itoa(wei%10) + res
		jinwei = wei / 10

	}
	return res
}

//func main() {
//	fmt.Println(addStrings("38", "65"))
//}
