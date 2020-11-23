package yahb

import (
	"errors"
	"fmt"
	"strconv"

	"encoding/json"
)

var (
	ErrInvalidSizeW  = errors.New("yahb: bad size width")
	ErrInvalidSizeH  = errors.New("yahb: bad size height")
	ErrSizeUnmarshal = errors.New("yahb: err size unmarshal")
)

type Size struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

// Validates the `size` object
func (s *Size) Validate() error {

	if s.Height == 0 {
		return ErrInvalidSizeH
	}

	if s.Width == 0 {
		return ErrInvalidSizeW
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Size) UnmarshalJSON(data []byte) (err error) {

	var stuff interface{}
	err = json.Unmarshal(data, &stuff)
	if err != nil {
		return ErrSizeUnmarshal
	}

	switch t := stuff.(type) {
	case []interface{}:

		for key, value := range t {
			if key == 0 {
				s.Width, err = _toInt(value)
				if err != nil {
					return ErrInvalidSizeW
				}
			} else if key == 1 {
				s.Height, err = _toInt(value)
				if err != nil {
					return ErrInvalidSizeH
				}
			}
		}

	case map[string]interface{}:
		for key, value := range t {
			if key == "width" {
				s.Width, err = _toInt(value)
				if err != nil {
					return ErrInvalidSizeW
				}
			} else if key == "height" {
				s.Height, err = _toInt(value)
				if err != nil {
					return ErrInvalidSizeH
				}
			}
		}
	}

	return nil
}

func _toInt(i interface{}) (int, error) {

	if i == nil {
		return 0, fmt.Errorf("empty value")
	}

	switch i.(type) {
	default:
		return 0, fmt.Errorf("bad value")
	case float32:
		return int(i.(float32)), nil
	case float64:
		return int(i.(float64)), nil
	case int:
		return i.(int), nil
	case int64:
		return int(i.(int64)), nil
	case string:
		return strconv.Atoi(i.(string))
	case json.Number:
		v, e := i.(json.Number).Int64()
		return int(v), e
	}
}
