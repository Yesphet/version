package version

import (
	"strconv"
	"encoding/json"
)

type Simple int64

func (v *Simple) String() string {
	return strconv.FormatInt(int64(*v), 10)
}

func (v *Simple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func (v *Simple) Release() {
	*v = Simple(int64(*v) + 1)
}

func newSimple() Version {
	v := Simple(0)
	return &v
}

func newSimpleFromString(raw string) (Version, error) {
	s, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return nil, err
	}

	v := Simple(s)

	return &v, nil
}
