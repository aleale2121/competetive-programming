package main

import "fmt"

func pivotIndex(nums []int) int {
	var prefix []int = []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		prefix = append(prefix, prefix[i-1]+nums[i])
	}
	if len(nums) == 1 {
		return 0
	}
	pivot := -1
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d  %d \n", prefix[i], prefix[len(nums)-1]-prefix[i])
		if i == 0 {
			if prefix[len(nums)-1]-prefix[i] == 0 {
				return i
			}
		} else if i == len(nums)-1 {
			if prefix[len(nums)-2] == 0 {
				return i
			}
		} else if prefix[i-1] == prefix[len(nums)-1]-prefix[i] {
			pivot = i
			break
		}
	}
	return pivot
}
