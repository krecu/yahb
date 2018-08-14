package yahb

import "errors"

var (
	ErrInvalidBidNoID         = errors.New("yahb: bid is missing ID")
	ErrInvalidBidNoCpm        = errors.New("yahb: bid is missing cpm")
	ErrInvalidBidNoCurr       = errors.New("yahb: bid is missing currency")
	ErrInvalidBidNoSize       = errors.New("yahb: bid is missing size")
	ErrInvalidBidNoDisplayUrl = errors.New("yahb: bid is missing DisplayUrl")
)

type Bid struct {
	ID          string    `json:"id"`                    // идентификатор для рекламного места в терминах Яндекса (из запроса)
	Cpm         float64   `json:"cpm"`                   // ставка, целое число, больше нуля
	Currency    string    `json:"currency,omitempty"`    // валюта ставки 'RUB'
	DisplayUrl  string    `json:"displayUrl,omitempty"`  // url за креативом или строка с кодом (html или js)
	DisplayCode string    `json:"displayCode,omitempty"` // url за креативом или строка с кодом (html или js)
	Size        *Size     `json:"size,omitempty"`
	Ext         Extension `json:"ext,omitempty"`
}

// Validate required attributes
func (bid *Bid) Validate() error {
	if bid.ID == "" {
		return ErrInvalidBidNoID
	} else if bid.Cpm <= 0 {
		return ErrInvalidBidNoCpm
	} else if bid.Currency == "" {
		return ErrInvalidBidNoCurr
	} else if bid.DisplayUrl == "" && bid.DisplayCode == "" {
		return ErrInvalidBidNoDisplayUrl
	} else if bid.Size == nil {
		return ErrInvalidBidNoSize
	}

	if err := bid.Size.Validate(); err != nil {
		return err
	}

	return nil
}
