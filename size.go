package yahb

import "errors"

var (
	ErrInvalidSizeW = errors.New("yahb: bad size width")
	ErrInvalidSizeH = errors.New("yahb: bad size height")
)

type Size struct {
	Width  int       `json:"width,omitempty"`
	Height int       `json:"height,omitempty"`
	Ext    Extension `json:"ext,omitempty"`
}

// Validates the `size` object
func (s *Size) Validate() error {
	if s.Width == 0 {
		return ErrInvalidSizeW
	} else if s.Height == 0 {
		return ErrInvalidSizeH
	}

	return nil
}
