package main

import (
	"fmt"
)

func compress(chars []byte) int {
	s := ""
	i := 0
	n := len(chars) - 1
	for i <= n {
		count := 1
		for j := n; j > i; j-- {
			if chars[j] == chars[i] {
				count++
			} else {
				count = 1
			}
		}
		s += string(chars[i])
		if count > 1 {
			s += fmt.Sprintf("%v", count)
			if i+count <= len(chars)-1 {
				i += count
			} else {
				break
			}
		} else {
			i++
		}
	}

	for j := 0; j < len(s); j++ {
		chars[j] = byte(s[j])
	}
	return len(s)
}
