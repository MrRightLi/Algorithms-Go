package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}

type Stack struct {
	data []int
	len int
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
	s.len++
}

func (s *Stack) Pop() (int,bool) {
	if s.len == 0 {
		return -1, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[0:len(s.data)-1]
	s.len--
	return top, true
}

func (s *Stack) Top() int {
	return s.data[len(s.data) - 1]
}

type MinStack struct {
	stack Stack
	min Stack
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.stack.Push(x)
	if this.min.len == 0 {
		this.min.Push(x)
	} else if x <= this.min.Top() {
		this.min.Push(x)
	}
}


func (this *MinStack) Pop()  {
	pop, r := this.stack.Pop()
	if r && pop == this.min.Top() {
		this.min.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.stack.Top()
}


func (this *MinStack) GetMin() int {
	return this.min.Top()
}


/**
 * Your MinStack object will be instantiated and called as such:
 * c
 * c
 * c
 * c
 * c
 */