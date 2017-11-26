// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package numutil_test

import (
	"testing"

	. "github.com/hoahm/myutil/numutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNumutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Numutil Suite")
}

var _ = Describe("Numutil", func() {
	Describe("Xrand", func() {
		It("should greater than or equal to start value", func() {
			Expect(Xrand(0, 100)).To(SatisfyAll(
				BeNumerically(">=", 0),
			))
		})

		It("should less than or equal to end value", func() {
			Expect(Xrand(0, 100)).To(SatisfyAll(
				BeNumerically("<=", 100),
			))
		})
	})

	Describe("AtoSI", func() {
		type myArray struct {
			s string
			a []int
		}

		It("splits string by space in to an array of integer", func() {
			for _, e := range []myArray{
				{ "0 1 2 3 4 5", []int{0, 1, 2, 3, 4, 5} },
				{ "9 3 6 1 2", []int{9, 3, 6, 1, 2} },
			} {
				Expect(AtoSI(e.s, " ")).To(Equal(e.a))
			}
		})

		It("splits string by comma in to an array of integer", func() {
			for _, e := range []myArray{
				{ "0,1,2,3,4,5", []int{0, 1, 2, 3, 4, 5} },
				{ "9,3,6,1,2", []int{9, 3, 6, 1, 2} },
			} {
				Expect(AtoSI(e.s, ",")).To(Equal(e.a))
			}
		})
	})

	Describe("IsDigit", func() {
		It("checks a string is a digit", func() {
			for _, s := range [...]string{
				"0", "001", "100", "-2147483648", "2147483647",
			} {
				Expect(IsDigit(s)).To(BeTrue())
			}
		})
		
		It("returns false if not a digit", func() {
			for _, s := range []string{
				"abc", "!@#$", "s100", "17A", "number1",
			} {
				Expect(IsDigit(s)).To(BeFalse())
			}
		})
	})
})
