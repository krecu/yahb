package yahb

import "errors"

var (
	ErrInvalidRespNoID       = errors.New("openrtb: response missing ID")
	ErrInvalidRespNoSeatBids = errors.New("openrtb: response missing seatbids")
)

// easyjson:json
type BidResponse struct {
	Bid []Bid `json:"bids,omitempty"` // Array of bid objects
}

// Validate attributes
func (res *BidResponse) Validate() error {

	if len(res.Bid) == 0 {
		return ErrInvalidRespNoSeatBids
	}

	for _, b := range res.Bid {
		if err := (&b).Validate(); err != nil {
			return err
		}
	}

	return nil
}
