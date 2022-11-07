package main

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums), len(nums))
	prefix := make([]int, len(nums), len(nums))
	suffix := make([]int, len(nums), len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	suffix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		ans[i] = prefix[i] * suffix[i]
	}
	return ans
}
