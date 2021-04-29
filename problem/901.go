package problem

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := new(list.List)
	return &Stack{list: list}
}

func (s *Stack) Pop() interface{} {
	first := s.list.Back()
	lst := s.list
	return lst.Remove(first)
}

func (s *Stack) Push(value interface{}) {
	lst := s.list
	lst.PushBack(value)
}

func (s *Stack) Peek() interface{} {
	return s.list.Back().Value
}

func (s *Stack) Len() int {
	return s.list.Len()
}

type StockSpanner struct {
	stack  *Stack
	weight *Stack
	list   []int
}

func Constructor() StockSpanner {
	s := new(StockSpanner)
	s.stack = NewStack()
	s.weight = NewStack()
	return *s
}

func (this *StockSpanner) Next(price int) int {
	this.list = append(this.list, price)
	index := len(this.list) - 1
	if this.stack.Len() == 0 {
		this.stack.Push(index)
		this.weight.Push(1)
		return 1
	}
	topIndex := this.stack.Peek().(int)
	if price < this.list[topIndex] {
		this.stack.Push(index)
		this.weight.Push(1)
		return 1
	}
	w := 1
	for this.stack.Len() > 0 && this.list[index] >= this.list[this.stack.Peek().(int)] {
		this.stack.Pop()
		w = w + this.weight.Pop().(int)
	}
	this.stack.Push(index)
	this.weight.Push(w)
	return w

}
