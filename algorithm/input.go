package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func p(row int, dp [][]int) () {
	for i := 0; i < row; i++ {
		for _, v := range dp[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println("")
	}
}

func know_row() {
	var row int
	fmt.Scanf("%d%d", &row)
	res := make([][]int, 0)
	for i := 0; i < row; i++ {
		input := bufio.NewReader(os.Stdin)
		str, _ := input.ReadString('\n')
		str = strings.Replace(str, "\n", "", -1)
		str_lis := strings.Split(str, " ")
		new_lis := make([]int, 0)
		for _, v := range str_lis {
			num, _ := strconv.Atoi(v)
			new_lis = append(new_lis, num)
		}
		res = append(res, new_lis)
	}
	fmt.Println("res=", res)
	p(row, res)

}

//func main() {
//	know_row()
//}
