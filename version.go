package version

import (
	"encoding/json"
	"fmt"
)

type Version interface {
	String() string
	json.Marshaler
}

func FromString(raw string) (Version, error) {

	v, err := newSimpleFromString(raw)
	if err == nil {
		return v, nil
	}

	v, err = newComplexFromString(raw)
	if err == nil {
		return v, nil
	}

	return nil, fmt.Errorf("unsupport version string %s", raw)
}
