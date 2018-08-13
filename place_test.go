package yahb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("place", func() {
	var subject *Place

	BeforeEach(func() {
		err := fixture("place", &subject)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should parse correctly", func() {
		Expect(subject).To(Equal(&Place{
			ID:          "1",
			PlacementId: "1",
			Sizes: []Size{
				Size{
					Width:  100,
					Height: 100,
				},
				Size{
					Width:  200,
					Height: 200,
				},
			},
		}))
	})

	It("should validate", func() {
		Expect((&Place{}).Validate()).To(Equal(ErrInvalidPlaceNoID))
	})

})
