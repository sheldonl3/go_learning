package main

import (
	"strconv"
)

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。
你可以假设字符串中只包含大小写英文字母（a至z）。
*/

func compressString(S string) string {
	l := len(S)
	if l < 2 {
		return S
	}
	res := ""
	for i := 0; i < l; i++ {
		tmp := 1
		if i == l-1 {
			res += string(S[i]) + strconv.Itoa(tmp)
			break
		}
		for S[i] == S[i+1] {
			tmp++
			i++
			if i >= l-1 {
				break
			}
		}
		res += string(S[i]) + strconv.Itoa(tmp)
	}
	if len(res) >= l {
		return S
	}
	return res
}
