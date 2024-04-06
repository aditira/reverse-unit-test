package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReverseString", func() {
	When("name is empty", func() {
		It("should return empty string", func() {
			Expect(ReverseString("")).To(Equal(""))
		})
	})

	When("name is not empty", func() {
		It("should return name in reverse order", func() {
			Expect(ReverseString("name")).To(Equal("eman"))
			Expect(ReverseString("aditira")).To(Equal("aritida"))
		})
	})

})
