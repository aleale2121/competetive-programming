package main

import "math"

func longestOnes(nums []int, k int) int {
	left, right, max := 0, 0, 0
	var ans float64 = 0
	for right < len(nums) {
		if nums[right] == 0 {
			max++
		}
		for max > k {
			if nums[left] == 0 {
				max--
			}
			left++
		}
		ans = math.Max(float64(ans), float64(right-left+1))
		right++
	}
	return int(ans)
}
