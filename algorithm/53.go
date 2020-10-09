package main

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray(nums []int) int {
	l := len(nums)
	max := INT_MIN
	sum := INT_MIN
	for i := 0; i < l; i++ { //sum小于0，更新计算；大于0，继续累加
		if sum < 0 {
			sum = nums[i]
			if max < sum {
				max = sum
			}
		} else {
			sum += nums[i]
			if max < sum {
				max = sum
			}
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
