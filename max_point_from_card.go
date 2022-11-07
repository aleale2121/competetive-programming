package main

func maxScore(cardPoints []int, k int) int {
	tot := findArraySum(cardPoints)
	n := len(cardPoints)
	curr := findArraySum(cardPoints[:n-k])
	ans := tot - curr
	for i := 0; i < k; i++ {
		curr -= cardPoints[i]
		curr += cardPoints[i+n-k]
		ans = max(ans, tot-curr)
	}
	return ans
}

func findArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
