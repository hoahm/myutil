package list_test

import (
	"testing"

	. "github.com/hoahm/myutil/datastructures/list"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "List Suit")
}

var _ = Describe("List", func() {
	Context("empty", func() {
		var l = New()

		It("should be empty", func() {
			Expect(l.IsEmpty()).To(BeTrue())
		})

		It("length is 0", func() {
			Expect(l.Len()).To(Equal(0))
		})
	})

	Context("not empty", func() {

	})
})