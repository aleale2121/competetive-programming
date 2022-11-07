package main

import "fmt"

func isValid(s string) bool {
	stk := new(Stack)
	for _, ch := range s {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}


type Stack struct {
	s []rune
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value rune) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() rune {
	if s.IsEmpty() == true {
		return ' '
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}
