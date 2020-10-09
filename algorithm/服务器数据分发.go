//我们共有n台服务器，每台服务器可以和若干个子服务器传输数据，n台服务器组成一个树状结构。
//现在要将一份数据从root节点开始分发给所有服务器。
//一次数据传输需要一个小时时间，
//一个节点可以同时对k个儿子节点进行并行传输，
//不同节点可以并行分发。
//问，全部分发完成，最短需要多少小时？
package main

import (
	"fmt"
	"sort"
)

var speed, row int
var data [][]int

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func get_time(row int) int {
	res := 0
	if row >= len(data) {
		return res
	}
	l := len(data[row])

	time_for_child := make([]int, 0)
	for i := 0; i < l; i++ {
		time_for_child = append(time_for_child, get_time(data[row][i]))
	}
	sort.Ints(time_for_child)
	time_for_child = reverse(time_for_child)
	i := 0
	for i < len(time_for_child) {
		time_for_child[i] = time_for_child[i] + i/speed + 1 // i/speed:排队的时间
		res = max(res, time_for_child[i])
		i++
	}
	return res
}

func input_data() {
	fmt.Scanf("%d%d", &speed, &row)
	data = make([][]int, 2000000)
	for i := 0; i < row; i++ {
		var num, partent int
		fmt.Scanf("%d%d", &num, &partent)
		data[partent] = make([]int, 0)
		for j := 0; j < num-1; j++ {
			var child int
			fmt.Scanf("%d", &child)
			data[partent] = append(data[partent], child)
		}
	}
	fmt.Println("data=", data)
	return

}
//
//func main() {
//	input_data()
//	fmt.Println(get_time(0))
//}
