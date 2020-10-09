package main

import "fmt"

func min(a,b,c int)int{
	if a<=b && a<=c{
		return a
	}else if b<=a && b<=c{
		return b
	}else {
		return c
	}
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i:=0;i<m+1;i++{
		dp[i][0]=i
	}
	for j:=0;j<n+1;j++{
		dp[0][j]=j
	}
	fmt.Println(dp)
	for j:=1;j<n+1;j++{
		for i:=1;i<m+1;i++{
			if word1[i-1]==word2[j-1]{
				dp[i][j]=dp[i-1][j-1]
			}else {
				dp[i][j]=min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

//func main() {
//	var n1 string
//	var n2 string
//	fmt.Scanln(&n1)
//	fmt.Scanln(&n2)
//	res := minDistance(n1, n2)
//	fmt.Println(res)
//}
