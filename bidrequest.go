package yahb

import "errors"

var (
	ErrInvalidReqNoPlaces   = errors.New("yahb: request has no places")
	ErrInvalidReqNoSettings = errors.New("yahb: request missing setting")
)

// easyjson:json
type BidRequest struct {
	Places  []Place  `json:"places,omitempty"`
	Setting *Setting `json:"settings,omitempty"`
}

// Validate attributes
func (req *BidRequest) Validate() error {
	if len(req.Places) == 0 {
		return ErrInvalidReqNoPlaces
	} else if req.Setting == nil {
		return ErrInvalidReqNoSettings
	}

	// validate settings
	if err := req.Setting.Validate(); err != nil {
		return err
	}

	// validate placements
	for _, p := range req.Places {
		if err := (&p).Validate(); err != nil {
			return err
		}
	}

	return nil
}
