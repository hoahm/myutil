// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package heap_test

import (
	"testing"

	. "github.com/hoahm/myutil/datastructures/heap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHeap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heap Suit")
}

var _ = Describe("Heap", func() {
	var (
		a = []int{7, 12, 6, 10, 17, 15, 2, 4}
	)

	Describe("empty", func() {
		var h Heap

		It("Len() returns 0", func() {
			Expect(h.Len()).To(Equal(0))
		})

		It("Size() returns 0", func() {
			Expect(h.Size()).To(Equal(0))
		})

		It("IsEmpty() returns true", func() {
			Expect(h.IsEmpty()).To(BeTrue())
		})
	})

	Describe("Build()", func() {
		h := Build(a)
		
		It("not empty", func() {
			Expect(h.IsEmpty()).To(BeFalse())
		})

		It("returns length of heap", func() {
			Expect(h.Len()).To(Equal(len(a)))
		})
		
		It("build a min heap", func() {
			Expect(h.Print()).To(Equal("2 4 6 10 17 15 7 12"))
		})

		It("returns smallest element", func() {
			Expect(h.Top()).To(Equal(2))
		})
	})

	Describe("Push()", func() {
		h := Build(a)
		length := len(a)
		h.Push(3)

		It("inserts new element", func() {	
			Expect(h.Print()).To(Equal("2 3 6 4 17 15 7 12 10"))
		})

		It("increase length", func() {
			Expect(h.Len()).To(Equal(length + 1))
		})
	})

	Describe("Pop()", func() {
		h := Build(a)
		length := h.Len()
		min, _ := h.Pop()

		It("removes first element", func() {
			Expect(min).To(Equal(2))
		})

		It("descrease the length", func() {
			Expect(h.Len()).To(Equal(length-1))
		})

		It("re-balance the heap", func() {
			Expect(h.Print()).To(Equal("4 10 6 12 17 15 7"))
		})
	})

	Describe("BuildMax()", func() {
		h := BuildMax(a)

		It("not empty", func() {
			Expect(h.IsEmpty()).To(BeFalse())
		})

		It("returns length of heap", func() {
			Expect(h.Len()).To(Equal(len(a)))
		})

		It("build a max heap", func() {
			Expect(h.Print()).To(Equal("17 12 15 10 7 6 2 4"))
		})

		It("returns largest element", func() {
			Expect(h.Top()).To(Equal(17))
		})
	})

	Describe("PushMax()", func() {
		h := BuildMax(a)
		h.PushMax(3)

		It("inserts new element", func() {
			Expect(h.Print()).To(Equal("17 12 15 10 7 6 2 4 3"))
		})

		It("increase length", func() {
			Expect(h.Len()).To(Equal(len(a) + 1))
		})
	})

	Describe("PopMax()", func() {
		h := BuildMax(a)
		length := h.Len()
		max, _ := h.PopMax()

		It("removes first element", func() {
			Expect(max).To(Equal(17))
		})

		It("descrease length", func() {
			Expect(h.Len()).To(Equal(length - 1))
		})

		It("re-balance the heap", func() {
			Expect(h.Print()).To(Equal("15 12 6 10 7 4 2"))
		})
	})
})