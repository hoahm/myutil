// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 

package fileutil_test

import (
	"testing"

	. "github.com/hoahm/myutil/fileutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFileutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fileutil Suite")
}

var _ = Describe("Fileutil", func() {
	Describe("CurrentDir", func() {
		It("returns string", func() {
			Expect(CurrentDir()).NotTo(BeEmpty())
		})
	})

	Describe("HomeDir", func() {
		It("returns string", func() {
			Expect(HomeDir()).NotTo(BeEmpty())
		})
	})

	Describe("Exists", func() {
		It("returns string", func() {
			Expect(Exists("./fileutil_test.go")).To(BeTrue())
		})
	})

	Describe("FilePath", func() {
		It("returns string", func() {
			Expect(FilePath("./", "fileutil_test", "go")).To(Equal("fileutil_test.go"))
		})
	})
})
