package yahb

import "errors"

var (
	ErrInvalidSettingCurr = errors.New("yahb: Settings missing currency")
)

type Setting struct {
	Currency string    `json:"currency,omitempty"`
	Ext      Extension `json:"ext,omitempty"`
}

// Validates the `settings` object
func (setting *Setting) Validate() error {
	if setting.Currency == "" {
		return ErrInvalidSettingCurr
	}

	return nil
}
