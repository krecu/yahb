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
		}))
	})

	It("should validate", func() {
		Expect((&Place{}).Validate()).To(Equal(ErrInvalidPlaceNoID))
	})

})
