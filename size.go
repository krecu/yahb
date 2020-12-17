package yahb

import (
	"errors"
)

var (
	ErrInvalidSizeW = errors.New("yahb: bad size width")
	ErrInvalidSizeH = errors.New("yahb: bad size height")
)

// easyjson:json
type Size struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

// Validate attributes
func (s *Size) Validate() error {

	if s.Height == 0 {
		return ErrInvalidSizeH
	}

	if s.Width == 0 {
		return ErrInvalidSizeW
	}

	return nil
}
