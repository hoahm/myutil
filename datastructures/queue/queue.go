package queue

type node struct {
	value interface{}
	next *node
}

// Queue is a first in first out (FIFO) data structure.
type Queue struct {
	head, tail *node
	length int
}

// New initialize an empty queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// IsEmpty check if queue is empty or not
func (q *Queue) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

// Len returns the number of elements
func (q *Queue) Len() int {
	return q.length
}

// Size is the same with Len()
func (q *Queue) Size() int {
	return q.length
}

// Front returns the 1st element of queue
func (q *Queue) Front() interface{} {
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
	temp := q.head
	q.head = temp.next
	q.length--
	
	return temp.value
}

// DeQueue is the same to Pop()
func (q *Queue) DeQueue() interface{} {
	return q.Pop()
}

// Push add new element to the end of queue
func (q *Queue) Push(value interface{}) {
	temp := &node{value, nil}
	if q.head == nil && q.tail == nil {
		q.head = temp 
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.length++
}

// EnQueue is the same to Push()
func (q *Queue) EnQueue(value interface{}) {
	q.EnQueue(value)
}