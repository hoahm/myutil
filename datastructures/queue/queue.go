// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package queue

import "sync"

type node struct {
	value interface{}
	next *node
}

// Queue is a first in first out (FIFO) data structure.
type Queue struct {
	head, tail *node
	length int
	sync.RWMutex
}

// New initialize an empty queue
func New() *Queue {
	return &Queue{nil, nil, 0, sync.RWMutex{}}
}

// IsEmpty check if queue is empty or not
func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	if q.length == 0 {
		return true
	}
	return false
}

// Len returns the number of elements
func (q *Queue) Len() int {
	q.RLock()
	defer q.RUnlock()

	return q.length
}

// Size is the same with Len()
func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return q.length
}

// Front returns the 1st element of queue
func (q *Queue) Front() interface{} {
	q.RLock()
	defer q.RUnlock()
	
	if q.head == nil {
		return nil
	}
	return q.head.value
}

// Pop get and remove 1st element of queue
func (q *Queue) Pop() interface{} {
	if q.head == nil {
		return nil
	}
	q.Lock()
	temp := q.head
	q.head = temp.next
	q.length--
	q.Unlock()
	return temp.value
}

// DeQueue is the same to Pop()
func (q *Queue) DeQueue() interface{} {
	return q.Pop()
}

// Push add new element to the end of queue
func (q *Queue) Push(value interface{}) {
	temp := &node{value, nil}
	q.Lock()
	if q.head == nil && q.tail == nil {
		q.head = temp 
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.length++
	q.Unlock()
}

// EnQueue is the same to Push()
func (q *Queue) EnQueue(value interface{}) {
	q.EnQueue(value)
}