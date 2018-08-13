package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BidRequest", func() {
	var subject *BidResponse

	BeforeEach(func() {
		err := fixture("bres", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&BidResponse{
			Bid: []Bid{Bid{
				ID:         "1",
				Currency:   "RUB",
				Cpm:        0.1,
				DisplayUrl: "http://test.com",
				Size: &Size{
					Width:  100,
					Height: 100,
				},
			}},
		}))
	})

	It("should validate", func() {
		Expect((&BidResponse{}).Validate()).To(Equal(ErrInvalidRespNoSeatBids))
		Expect(subject.Validate()).NotTo(HaveOccurred())
	})

})
