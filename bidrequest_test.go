package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidRequest", func() {
	var subject *BidRequest

	BeforeEach(func() {
		err := fixture("breq", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&BidRequest{
			Places: []Place{Place{
				ID:          "1",
				PlacementId: "1",
			}},
			Setting: &Setting{
				Currency: "RUB",
			},
		}))
	})

	It("should validate", func() {
		Expect((&BidRequest{}).Validate()).To(Equal(ErrInvalidReqNoPlaces))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
