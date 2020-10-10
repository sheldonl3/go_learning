package main

import (
	"fmt"
	"math/rand"
	"time"
)

//index区间0 - len-i生成随机数，将此位置的数选出与i位置的交换
func Shuffle(array []int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := len(array) - 1; i >= 0; i-- {
		p := RandInt64(0, int64(i))
		fmt.Printf("每次生成的随机数：%d\n", p)
		a := array[i]
		array[i] = array[p]
		array[p] = a
		fmt.Println("置换后的数组为：", array)
	}
	return array
}

// RandInt64 区间随机数
func RandInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
