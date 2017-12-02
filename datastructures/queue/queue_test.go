// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package queue_test

import (
	"testing"

	. "github.com/hoahm/myutil/datastructures/queue"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Suit")
}

var _ = Describe("Queue", func() {
	Context("empty", func() {
		var q = New()

		It("should be empty", func() {
			Expect(q.IsEmpty()).To(BeTrue())
		})

		It("length must be 0", func() {
			Expect(q.Len()).To(Equal(0))
		})

		It("head element is nil", func() {
			Expect(q.Front()).To(BeNil())
		})

		It("pop return nil", func() {
			Expect(q.Pop()).To(BeNil())
		})
	})

	Context("not empty", func() {
		var q Queue

		BeforeSuite(func() {
			q.Push(1)
			q.Push(2)
		})

		It("should not be empty", func() {
			Expect(q.IsEmpty()).To(BeFalse())
		})

		It("length must be 2", func() {
			Expect(q.Len()).To(Equal(2))
		})

		It("head element is 1", func() {
			Expect(q.Front()).To(Equal(1))
		})

		It("pop return head element", func() {
			head := q.Front()
			Expect(q.Pop()).To(Equal(head))
		})

		It("push increase length", func() {
			length := q.Len()
			q.Push(3)
			Expect(q.Len()).To(Equal(length + 1))
		})
	})
}) 