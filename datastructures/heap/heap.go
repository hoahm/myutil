// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package heap

import (
	// "fmt"
	"errors"
	"strconv"
)

// Heap is a data structure
type Heap struct {
	a []int
	length int
}

func build(a []int) *Heap {
	b := append([]int(nil), a...)
	copy(b, a)
	return &Heap{b, len(a)}
}

// Len returns the length of heap
func (h *Heap) Len() int {
	return h.length
}

// Size returns the length of heap
func (h *Heap) Size() int {
	return h.Len()
}

// IsEmpty check if heap is empty or not
func (h *Heap) IsEmpty() bool {
	if h.length == 0 {
		return true
	}
	return false
}

// minHeapify build a min heap
func (h *Heap) minHeapify(i int) {
	var (
		smallest = i
		left = 2 * i + 1
		right = 2 * i + 2
	)

	if left < h.length && h.a[left] < h.a[smallest] {
		smallest = left
	}

	if right < h.length && h.a[right] < h.a[smallest] {
		smallest = right
	}

	if smallest != i {		
		h.a[i], h.a[smallest] = h.a[smallest], h.a[i]

		h.minHeapify(smallest)
	}
}

// maxHeapify build a max heap
func (h *Heap) maxHeapify(i int) {
	var (
		largest = i
		left = 2 * i + 1
		right = 2 * i + 2
	)

	if left < h.length && h.a[left] > h.a[largest] {
		largest = left
	}

	if right < h.length && h.a[right] > h.a[largest] {
		largest = right
	}

	if largest != i {
		h.a[i], h.a[largest] = h.a[largest], h.a[i]

		h.maxHeapify(largest)
	}
}

// Build Heap
func Build(a []int) *Heap {
	h := build(a)

	for i := h.length / 2 - 1; i >= 0; i-- {
		h.minHeapify(i)
	} 

	return h
}

// BuildMax heap
func BuildMax(a []int) *Heap {
	h := build(a)

	for i:= h.length / 2 - 1; i >= 0; i-- {
		h.maxHeapify(i)
	}

	return h
}

// Print export heap to string
func (h *Heap) Print() string {
	str := ""
	for i := 0; i < h.length; i++ {
		str += strconv.Itoa(h.a[i])
		if i != h.length - 1 {
			str += " "
		}
	}
	return str
}

// Top returns the 1st element of the heap
// If min heap, it is the smallest
// If max heap, it is the largest
func (h *Heap) Top() (int, error) {
	if h.length == 0 {
		return 0, errors.New("Heap is empty")
	}
	return h.a[0], nil
}

// Push insert new value to min heap
func (h *Heap) Push(val int) {
	h.a = append(h.a, val)
	h.length++

	i := h.length - 1
	for i != 0 && h.a[(i - 1) / 2] > h.a[i] {
		h.a[i], h.a[(i - 1) / 2] = h.a[(i - 1) / 2], h.a[i]
		i = (i - 1) / 2
	}
}

// PushMax insert new value to max heap 
func (h *Heap) PushMax(val int) {
	h.a = append(h.a, val)
	h.length++

	i := h.length - 1
	for i != 0 && h.a[(i - 1) / 2] < h.a[i] {
		h.a[i], h.a[(i - 1) / 2] = h.a[(i - 1) / 2], h.a[i]
		i = (i - 1) / 2
	}
}

// Pop removes the minimum element and returns it
func (h *Heap) Pop() (int, error) {
	if h.length == 0 {
		return 0, errors.New("Heap is empty")
	}

	// Get the 1st element
	min := h.a[0]

	// Assign last element to the the 1st element
	h.a[0] = h.a[h.length - 1]

	// Remove the last element
	h.a = h.a[:h.length - 1]
	h.length--

	// Re-balance heap
	h.minHeapify(0)

	// fmt.Println(h.a)

	return min, nil
}

// PopMax removes the maximum element and returns it
func (h *Heap) PopMax() (int, error) {
	if h.length == 0 {
		return 0, errors.New("Heap is empty")
	}

	// Get the 1st element
	max := h.a[0]

	// Assign last element to the 1st element
	h.a[0] = h.a[h.length - 1]

	// Remove the last element
	h.a = h.a[:h.length - 1]
	h.length--

	// Re-balance the heap
	h.maxHeapify(0)

	return max, nil
}

func (h *Heap) Expose() []int {
	return h.a
}