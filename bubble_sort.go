package main

import (
	"fmt"
)

func countSwaps(a []int32) {
	isdone := false
	swap := 0
	for !isdone {
		isdone = true
		i := 0
		for i < len(a)-1 {
			if a[i+1] < a[i] {
				a[i], a[i+1] = a[i+1], a[i]
				isdone = false
				swap++
			}
			i++
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", swap)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])
}
