package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bid", func() {
	var subject *Bid

	BeforeEach(func() {
		err := fixture("bid", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Bid{
			ID:         "1",
			Currency:   "RUB",
			Cpm:        0.1,
			DisplayUrl: "http://test.com",
			Size: &Size{
				Width:  100,
				Height: 100,
			},
		}))
	})

	It("should validate", func() {
		Expect((&Bid{}).Validate()).To(Equal(ErrInvalidBidNoID))
	})

})
