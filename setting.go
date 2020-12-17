package yahb

import "errors"

var (
	ErrInvalidSettingCurr = errors.New("yahb: settings missing currency")
)

// easyjson:json
type Setting struct {
	Currency string `json:"currency,omitempty"`
	WinSize  *Size  `json:"windowSize,omitempty"`
}

// Validate required attributes
func (setting *Setting) Validate() error {

	if setting.Currency == "" {
		return ErrInvalidSettingCurr
	}

	if setting.WinSize != nil {
		if err := setting.WinSize.Validate(); err != nil {
			return err
		}
	}

	return nil
}
