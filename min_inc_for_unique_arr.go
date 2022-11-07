package main

import (
	"math"
	"sort"
)

func minIncrementForUnique(nums []int) int {
	var ans float64 = 0
	var minAvailable float64 = 0
	sort.Ints(nums)
	for _, v := range nums {
		ans += math.Max(minAvailable-float64(v), 0)
		minAvailable = math.Max(minAvailable, float64(v)) + 1
	}
	return int(ans)
}
