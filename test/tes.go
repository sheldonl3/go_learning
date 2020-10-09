package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			//fmt.Println(res)
			if nums[l]+nums[r]+nums[i] == 0 {
				var tmp = []int{nums[l], nums[r], nums[i]}
				res = append(res, tmp)
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				r -= 1
				l += 1
			} else if nums[l]+nums[r]+nums[i] > 0 {
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSum(nums)
}
