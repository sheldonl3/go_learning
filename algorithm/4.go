package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lis := make([]int, 0)
	lis = append(lis, nums1...)
	lis = append(lis, nums2...)
	sort.Ints(lis)
	l:=len(lis)
	mid:=len(lis)/2
	if l%2==0{
		return float64((lis[mid-1]+lis[mid]))/2
	}
	return float64(lis[mid])
}

func input() [][]int {
	res := make([][]int, 0)
	for i := 0; i < 2; i++ {
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
	return res
}

//func main() {
//	in := input()
//	fmt.Println(findMedianSortedArrays(in[0], in[1]))
//}
