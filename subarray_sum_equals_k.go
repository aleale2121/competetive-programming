package main

func subarraySum(nums []int, k int) int {
	var prefix = 0
	var count = map[int]int{0: 1}
	var ans = 0
	for _, v := range nums {
		prefix += v
		ans += count[prefix-k]
		count[prefix] += 1
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
