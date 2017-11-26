// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 
package stringutil_test

import (
	"testing"

	. "github.com/hoahm/myutil/stringutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringutil Suite")
}

var _ = Describe("Stringutil", func() {
	Describe("RandString", func() {
		It("does not return empty string", func() {
			Expect(RandString(6)).NotTo(BeEmpty())
		}) 

		It("does not return nil", func() {
			Expect(RandString(6)).NotTo(BeNil())
		})

		It("length is 6 characters", func() {
			Expect(RandString(6)).To(HaveLen(6))
		})
	})

	Describe("Reverse", func() {
		It("returns string in the reversed order", func() {
			for _, c := range []struct {
				in, want string
			}{
				{"Hello, world", "dlrow ,olleH"},
				{"Hello, 世界", "界世 ,olleH"},
				{"", ""},
			} {
				Expect(Reverse(c.in)).To(Equal(c.want))
			}
		})
	})
})
