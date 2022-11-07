package main

func numberOfSubarrays(nums []int, k int) int {
	subArraysAtmosK := func(n int) int {
		ans, left, right := 0, 0, 0
		for right <= len(nums) {
			if n >= 0 {
				ans += right - left
				if right == len(nums) {
					break
				}
				if nums[right]%2 == 1 {
					n--
				}
				right++
			} else {
				if nums[left]%2 == 1 {
					n++
				}
				left++
			}

		}
		return ans
	}
	return subArraysAtmosK(k) - subArraysAtmosK(k-1)
}
