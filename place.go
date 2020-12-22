package yahb

import "errors"

var (
	ErrInvalidPlaceNoID = errors.New("yahb: ID missing")
	ErrInvalidNoPlaceId = errors.New("yahb: placement ID missing")
)

// easyjson:json
type Place struct {
	ID          string      `json:"id"`                    // уникальный идентификатор Яндекса для рекламного места - должеа вернуться в ответе
	PlacementId string      `json:"placementId,omitempty"` // идентификатор рекламного в терминах монетизатора
	CodeType    BidCodeType `json:"codeType,omitempty"`
	Sizes       [][]int     `json:"sizes,omitempty"` // Массив размеров, которые указал пользователь [ширина,высота]
}

// Size - size in uniq format
func (place *Place) Size() []Size {
	if len(place.Sizes) == 0 {
		return nil
	}

	var sizes []Size
	for _, s := range place.Sizes {
		sizes = append(sizes, Size{
			Width:  s[0],
			Height: s[1],
		})
	}

	return sizes
}

// Validate attributes
func (place *Place) Validate() error {

	if place.ID == "" {
		return ErrInvalidPlaceNoID
	}

	if place.PlacementId == "" {
		return ErrInvalidNoPlaceId
	}

	if len(place.Sizes) != 0 {
		for _, s := range place.Sizes {
			if err := (&Size{
				Width:  s[0],
				Height: s[1],
			}).Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
