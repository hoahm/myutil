// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package stack

import "sync"

type node struct {
	value interface{}
	prev *node	
}

// Stack is a last in - first out (LIFO) data structure
type Stack struct {
	top *node
	length int
	sync.RWMutex
}

// New initialize an empty stack
func New() *Stack {
	return &Stack{nil, 0, sync.RWMutex{}}
}

// IsEmpty check if stack is empty or not
func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()

	if s.length == 0 {
		return true
	}
	return false
}

// Len returns the number elements of stack
func (s *Stack) Len() int { 
	s.RLock()
	defer s.RUnlock()

	return s.length 
}

// Size is the same with Len()
func (s *Stack) Size() int {
	s.RLock()
	defer s.RUnlock()

	return s.length
}

// Top get the value of the peak element of stack
func (s *Stack) Top() interface{} {
	s.RLock()
	defer s.RUnlock()
	
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop get and remove the value of the peak element of stack
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	s.Lock()
	top := s.top
	s.top = top.prev
	s.length--
	s.Unlock()
	return top.value
}

// Push add new element to the peak of stack
func (s *Stack) Push(value interface{}) {
	s.Lock()
	n := &node{value, s.top}
	s.top = n
	s.length++
	s.Unlock()
}