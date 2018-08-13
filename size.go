package yahb

import (
	"errors"

	"encoding/json"
)

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

// UnmarshalJSON implements json.Unmarshaler
func (s *Size) UnmarshalJSON(data []byte) (err error) {

	var stuff interface{}
	err = json.Unmarshal(data, &stuff)
	if err != nil {
		return err
	}
	switch t := stuff.(type) {
	case []interface{}:

		for key, value := range t {
			if key == 0 {
				s.Width = ToInt(value)
			} else if key == 1 {
				s.Height = ToInt(value)
			}
		}

	case map[string]interface{}:
		for key, value := range t {
			if key == "width" {
				s.Width = ToInt(value)
			} else if key == "height" {
				s.Height = ToInt(value)
			}
		}
	}

	return nil
}

func ToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch i2 := i.(type) {
	default:
		return 0
	case *json.Number:
		i3, _ := i2.Int64()
		return int(i3)
	case json.Number:
		i3, _ := i2.Int64()
		return int(i3)
	case int64:
		return int(i2)
	case float64:
		return int(i2)
	case float32:
		return int(i2)
	case uint64:
		return int(i2)
	case int:
		return int(i2)
	case uint:
		return int(i2)
	case bool:
		if i2 {
			return 1
		} else {
			return 0
		}
	case *bool:
		if i2 == nil {
			return 0
		}
		if *i2 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
