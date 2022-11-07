package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	var ans float64 = 0
	var counter = make(map[byte]int, 128)
	left := 0
	for right := 0; right < len(s); right++ {
		counter[s[right]]++
		for counter[s[right]] > 1 {
			counter[s[left]]--
			left++
		}
		ans = math.Max(ans, float64(right-left+1))
	}
	return int(ans)
}
