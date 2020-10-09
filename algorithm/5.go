package main

import "fmt"

func judent(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 || l == 1 {
		return s
	}
	for i := l; i > 0; i-- {
		for j := 0; j <= l-i; j++ {
			fmt.Println("i=", i, "j=", j, "s=", s[j:i])
			if judent(s[j:i+j]) == true {
				return s[j:i+j]
			}
		}
	}
	return ""
}

//func main() {
//	res := longestPalindrome("baa")
//	fmt.Println(res)
//}
