package version

import (
	"fmt"
	"encoding/json"
	"strings"
	"strconv"
)

const separator = "."

type Complex struct {
	major int64
	minor int64
	patch int64
}

func (v *Complex) String() string {
	return fmt.Sprintf("%d%s%d%s%d", v.major, separator, v.minor, separator, v.patch)
}

func (v *Complex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

func newComplex() Version {
	return &Complex{
		major: 0,
		minor: 0,
		patch: 0,
	}
}

func newComplexFromString(s string) (Version, error) {
	var major, minor, patch int64
	var err error

	raw := strings.Split(s, separator)

	if len(raw) != 3 {
		goto errorReturn
	}

	major, err = strconv.ParseInt(raw[0], 10, 32)
	if err != nil {
		goto errorReturn
	}

	minor, err = strconv.ParseInt(raw[1], 10, 32)
	if err != nil {
		goto errorReturn
	}

	patch, err = strconv.ParseInt(raw[2], 10, 32)
	if err != nil {
		goto errorReturn
	}

	return &Complex{
		major: major,
		minor: minor,
		patch: patch,
	}, nil

errorReturn:
	return &Complex{}, fmt.Errorf("illegal version string")

}

func (v *Complex) ReleasePatch() {
	v.patch ++
}

func (v *Complex) ReleaseMinor() {
	v.minor ++
	v.patch = 0
}

func (v *Complex) ReleaseMajor() {
	v.major ++
	v.minor = 0
	v.patch = 0
}
