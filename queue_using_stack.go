package main

import "fmt"

type MyQueue struct {
	s1 StackInt
	s2 StackInt
}

func Constructor() MyQueue {
	return *new(MyQueue)
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.s[len(this.s2.s)-1]
}

func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

type StackInt struct {
	s []int
}

func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

func (s *StackInt) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *StackInt) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}
