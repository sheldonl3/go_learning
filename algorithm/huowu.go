package main

import (
	"strconv"
)

func zhuangzai(str []string) int {
	if len(str) != 4 {
		return -1
	}
	a, err := strconv.Atoi(str[0])
	b, err := strconv.Atoi(str[1])
	k, err := strconv.Atoi(str[2])
	v, err := strconv.Atoi(str[3])
	if err != nil {
		return -1
	}
	var put int
	var res int
	for {
		if put < a {
			if b > 0 {
				res++
				if b >= k-1 {
					b -= k - 1
					put += k * v
				} else {
					put += (b + 1) * v
					b = -1
				}
			} else {
				res++
				put += v
			}
		} else {
			break
		}
	}
	return res
}

//func main() {
//	for {
//		input := ScanLine()
//		s := strings.Split(input, " ")
//		//fmt.Println(s)
//		fmt.Println(zhuangzai(s))
//	}
//}
