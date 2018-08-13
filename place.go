package yahb

import "errors"

var (
	ErrInvalidPlaceNoID = errors.New("yahb: ID missing")
	ErrInvalidNoPlaceId = errors.New("yahb: placement ID missing")
)

type Place struct {
	ID          string    `json:"id"`                    // уникальный идентификатор Яндекса для рекламного места - должеа вернуться в ответе
	PlacementId string    `json:"placementId,omitempty"` // идентификатор рекламного в терминах монетизатора
	Sizes       *Size     `json:"sizes,omitempty"`       // Массив размеров, которые указал пользователь [ширина,высота]
	Ext         Extension `json:"ext,omitempty"`         // вдруг расширят
}

// Validates the `places` object
func (place *Place) Validate() error {
	if place.ID == "" {
		return ErrInvalidPlaceNoID
	}

	if place.PlacementId == "" {
		return ErrInvalidNoPlaceId
	}

	return nil
}
