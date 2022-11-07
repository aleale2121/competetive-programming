package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	max := 0
	score := 0
	i := 0
	j := len(tokens) - 1
	sort.Ints(tokens)

	for i <= j && (score > 0 || power >= tokens[i]) {
		for i <= j && power >= tokens[i] {
			power -= tokens[i]
			i++
			score++
		}
		if score > max {
			max = score
		}
		if i <= j && score > 0 {
			power += tokens[j]
			j--
			score--
		}
	}
	return max
}
