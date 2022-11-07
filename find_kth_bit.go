package main

import "math"

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	noCol := int(math.Pow(float64(2), float64(n))) - 1
	mid := noCol / 2

	if k <= mid {
		return findKthBit(n-1, k)
	} else if k == mid+1 {
		return '1'
	} else {
		k = noCol - k + 1
		if findKthBit(n-1, k) == '0' {
			return '1'
		}
		return '0'
	}

}
