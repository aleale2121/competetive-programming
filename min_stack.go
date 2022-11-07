package main

import "math"

type MinStack struct {
    s []int
    min int
}


func Constructor() MinStack {
    return MinStack{min: math.MaxInt64}
}


func (this *MinStack) Push(val int)  {
    this.s = append(this.s, val)
    if val < this.min {
		this.min = val
	}
}


func (this *MinStack) Pop()  {
  length := len(this.s)
  last := this.s[length-1]
  this.s = this.s[:length-1] 
  if last != this.min {
    return
  }
  this.min = math.MaxInt64
  for _, x := range this.s {
    if x < this.min {
      this.min = x
    }
  }
}


func (this *MinStack) Top() int {
    if len(this.s) == 0 {
		return 0
	}
	length := len(this.s)
	res := this.s[length-1]
	return res
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */