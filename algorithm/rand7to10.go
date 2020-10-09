package main

import (
	"math/rand"
)

var seed = [][]int{{1, 2, 3, 4, 5, 6, 7},
	{8, 9, 10, 1, 2, 3, 4},
	{5, 6, 7, 8, 9, 10, 1},
	{2, 3, 4, 5, 6, 7, 8},
	{9, 10, 1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10, 0, 0},
	{0, 0, 0, 0, 0, 0, 0}}

func rand7() int {
	return rand.Intn(7)
}

func rand7to10() int {
	i := 0
	for i == 0 {
		x := rand7()
		y := rand7()
		i = seed[x][y]
	}
	return i
}

//func main() {
//	rand.Seed(time.Now().UnixNano())
//	fmt.Println("got ", rand7to10())
//}
