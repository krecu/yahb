package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Size", func() {
	var subject *Size

	BeforeEach(func() {
		err := fixture("size", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Size{
			Width:  100,
			Height: 100,
		}))
	})

	It("should validate", func() {
		Expect((&Size{}).Validate()).To(Equal(ErrInvalidSizeW))
	})

})
