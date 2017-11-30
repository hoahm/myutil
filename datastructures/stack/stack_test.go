// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package stack_test

import (
	"testing"

	. "github.com/hoahm/myutil/datastructures/stack"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suit")
}

var _ = Describe("Stack", func() {
	Context("empty", func() {
		var s = New()

		It("should be empty", func() {
			Expect(s.IsEmpty()).To(BeTrue())
		})

		It("length is 0", func() {
			Expect(s.Len()).To(Equal(0))
		})

		It("peak element is nil", func() {
			Expect(s.Top()).To(BeNil())
		})

		It("pop return nil", func() {
			Expect(s.Pop()).To(BeNil())
		})
	})

	Context("not empty", func() {
		var s Stack

		BeforeSuite(func() {
			s.Push(1)
		})

		It("should not be empty", func() {
			Expect(s.IsEmpty()).To(BeFalse())
		})

		It("length is 1", func() {
			Expect(s.Len()).To(Equal(1))
		})

		It("push increase len", func() {
			s.Push(2)
			Expect(s.Len()).To(Equal(2))
		})

		It("peak element is not nil", func() {
			Expect(s.Top()).NotTo(BeNil())
		})

		It("peak element has value 2", func() {
			Expect(s.Top()).To(Equal(2))
		})

		It("pop return 2", func() {
			Expect(s.Pop()).To(Equal(2))
		})
	})
})