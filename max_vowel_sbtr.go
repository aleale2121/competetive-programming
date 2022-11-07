package main

func maxVowels(s string, k int) int {
	curr := vowelCount(s[0:k])
	ans := curr
	for r := k; r < len(s); r++ {
		curr += vowelCount(s[r:r+1]) - vowelCount(s[r-k:r-k+1])
		ans = max(ans, curr)
	}

	return ans
}

func vowelCount(s string) int {
	count := 0
	for _, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
