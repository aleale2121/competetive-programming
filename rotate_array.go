package main

func rotate(nums []int, k int) {
	var temp = make([]int, len(nums))
	m := 0
	k = k % len(nums)
	for i := len(nums) - k; i < len(nums); i++ {
		temp[m] = nums[i]
		m++
	}
	for j := 0; j < len(nums)-k; j++ {
		temp[m] = nums[j]
		m++
	}
	for j := 0; j < len(nums); j++ {
		nums[j] = temp[j]
	}

}
