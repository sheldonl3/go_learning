package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func make_heap(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	for i := l / 2; i >= 0; i-- {
		judge(nums, i,l)
	}
}

func judge(nums []int, index ,l int) {
	left := index*2 + 1
	right := index*2 + 2
	ma := 0
	max_num := 0
	if right < l {
		max_num = max(nums[left], nums[right])
		if max_num == nums[left] {
			ma = left
		} else {
			ma = right
		}
	} else if left < l {
		max_num = nums[left]
		ma = left
	} else {
		return
	}
	if max_num > nums[index] {
		nums[ma], nums[index] = nums[index], nums[ma]
		judge(nums, ma,l)
	}
}

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	if l < k {
		return -1
	}
	make_heap(nums)
	fmt.Println(nums)
	res:=0
	in:=l
	for i:=0;i<k;i++{
		res=nums[0]
		nums[0],nums[l-1-i]=nums[l-1-i],nums[0]
		in-=1
		judge(nums,0,in)
	}
	fmt.Println(nums)
	return res
}

func main() {
	var nums = []int{1,2}
	fmt.Println(findKthLargest(nums, 2))
	return
}
