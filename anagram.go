package main

func findAnagrams(s string, p string) []int {
	var count = make(map[rune]int, 128)
	var ans []int
	length := len(p)
	for _, v := range p {
		count[v]++
	}

	for r, v := range s {
		count[v]--
		if count[v] >= 0 {
			length--
		}
		if r >= len(p) {
			count[rune(s[r-len(p)])] += 1
			if count[rune(s[r-len(p)])] > 0 {
				length += 1
			}
		}
		if length == 0 {
			ans = append(ans, r-len(p)+1)
		}
	}
	return ans
}
