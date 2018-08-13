package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setting", func() {
	var subject *Setting

	BeforeEach(func() {
		err := fixture("setting", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Setting{
			Currency: "RUB",
		}))
	})

	It("should validate", func() {
		Expect((&Setting{}).Validate()).To(Equal(ErrInvalidSettingCurr))
	})
})
