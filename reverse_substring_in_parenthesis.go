package main

func reverseParentheses(s string) string {
	var stack [][]byte
	var top []byte
	stack = append(stack, []byte{}) 
	for i := 0; i < len(s); i++ {
	  if s[i] == '(' { 
		stack = append(stack, []byte{})
	  } else if s[i] == ')' {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		reverse(top)
		stack[len(stack)-1] = append(stack[len(stack)-1], top...)
	  } else { 
		stack[len(stack)-1] = append(stack[len(stack)-1], s[i])
	  }
	}
	return string(stack[0])
  }
  
  func reverse(seg []byte) {
	i, j := 0, len(seg)-1
	for i < j {
	  seg[i], seg[j] = seg[j], seg[i]
	  i++
	  j--
	}
  }