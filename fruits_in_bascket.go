package main

func totalFruit(fruits []int) int {

	ans := 0
	left := 0
	counter := make(map[int]int)
	for right := 0; right < len(fruits); right++ {
		counter[fruits[right]]++
		for len(counter) > 2 {
			counter[fruits[left]]--
			if counter[fruits[left]] == 0 {
				delete(counter, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
