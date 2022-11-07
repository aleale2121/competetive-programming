package main

func minSquare(n, m, a int64) int64 {
	c, b := n/a, m/a
	if n%a != 0 {
		c++
	}
	if m%a != 0 {
		b++
	}
	return c * b
}
