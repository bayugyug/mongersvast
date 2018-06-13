package mongersvast

import "errors"

// HashExt is a raw encoded JSON value.
type HashExt []byte

// MarshalJSON returns e as the JSON encoding of e.
func (e HashExt) MarshalJSON() ([]byte, error) {
	return e, nil
}

// UnmarshalJSON sets *e to a copy of data.
func (e *HashExt) UnmarshalJSON(data []byte) error {
	if e == nil {
		return errors.New("HashExt: UnmarshalJSON on nil pointer")
	}
	*e = append((*e)[0:0], data...)
	return nil
}
