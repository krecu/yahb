package yahb

import "errors"

var (
	ErrInvalidPlaceNoID = errors.New("yahb: ID missing")
	ErrInvalidNoPlaceId = errors.New("yahb: placement ID missing")
)

// easyjson:json
type Place struct {
	ID          string `json:"id"`                    // уникальный идентификатор Яндекса для рекламного места - должеа вернуться в ответе
	PlacementId string `json:"placementId,omitempty"` // идентификатор рекламного в терминах монетизатора
	Sizes       []Size `json:"sizes,omitempty"`       // Массив размеров, которые указал пользователь [ширина,высота]
}

// Validates the `places` object
func (place *Place) Validate() error {

	if place.ID == "" {
		return ErrInvalidPlaceNoID
	}

	if place.PlacementId == "" {
		return ErrInvalidNoPlaceId
	}

	if len(place.Sizes) != 0 {
		for _, s := range place.Sizes {
			if err := (&s).Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
