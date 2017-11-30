package list

import "sync"

type node struct {
	value interface{}
	prev, next *node
}

// List is a double linked list implement in Golang
// 
// Golang has their built-in linked list as describe at
// https://golang.org/pkg/container/list/
// So use this library as a reference for practicing
//
type List struct {
	head, tail *node
	length int
	sync.RWMutex
}

// New initialize an empty linked list
func New() *List {
	return &List{nil, nil, 0, sync.RWMutex{}}
}

// IsEmpty checks if linked list is empty or not
func (l *List) IsEmpty() bool {
	l.RLock()
	defer l.RUnlock()

	if l.length == 0 {
		return true
	}
	return false
}

// Len returns the number of elements
func (l *List) Len() int {
	l.RLock()
	defer l.RUnlock()

	return l.length
}

// Size is the same with Len()
func (l *List) Size() int {
	return l.Len()
}

// Front returns the head element
func (l *List) Front() interface{} {
	l.RLock()
	defer l.RUnlock()

	if l.length == 0 {
		return nil
	}
	return l.head.value
}

// Tail returns the tail element
func (l *List) Tail() interface{} {
	l.RLock()
	defer l.RUnlock()

	if l.length == 0 {
		return nil
	}
	return l.tail.value
}

// PushBack insert element to the tail
func (l *List) PushBack(value interface{}) {
	temp := &node{value, nil, nil}
	l.Lock()
	if l.length == 0 {
		l.head = temp
		l.tail = temp
	} else {
		temp.prev = l.tail
		l.tail = temp
	}
	l.length++
	l.Unlock()
}

// PushFront insert element to the head
func (l *List) PushFront(value interface{}) {
	temp := &node{value, nil, nil}
	l.Lock()
	if l.length == 0 {
		l.head = temp
		l.tail = temp
	} else {
		temp.next = l.head
		l.head = temp
	}
	l.length++
	l.Unlock()
}

// Remove finds and removes element from list
func (l *List) Remove(value interface{}) {
	l.Lock()
	for n := l.head; n != nil; n = n.next {
		if n.value == value {
			n.prev.next = n.next
			n.next.prev = n.prev
			l.length--
		}
	}
	l.Unlock()
}

// Merge 2 linked list together
func (l *List) Merge(l2 *List) {
	if l.length == 0 {
		l.Lock()
		l2.Lock()
		l = l2
		l.Unlock()
		l2.Unlock()
		return
	}

	if l2.length == 0 {
		return
	}
	l.Lock()
	l2.Lock()
	l.tail.next = l2.head
	l2.head.prev = l.tail
	l.tail = l2.tail
	l.length += l2.length
	l.Unlock()
	l2.Unlock()
}