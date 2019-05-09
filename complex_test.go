package version

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"fmt"
)

func TestComplex_String(t *testing.T) {
	v, err := newComplexFromString("1.2.3")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", v.String())

	v, err = newComplexFromString("1.2.3")
	assert.Nil(t, err)
	assert.Equal(t, "1.2.3", v.String())
}

func TestComplex_ReleaseMajor(t *testing.T) {
	v, _ := newComplexFromString("1.2.3")
	v.(*Complex).ReleaseMajor()
	assert.Equal(t, "2.0.0", v.String())
}

func TestComplex_ReleaseMinor(t *testing.T) {
	v, _ := newComplexFromString("1.2.3")
	v.(*Complex).ReleaseMinor()
	assert.Equal(t, "1.3.0", v.String())
}

func TestComplex_ReleasePatch(t *testing.T) {
	v, _ := newComplexFromString("1.2.3")
	v.(*Complex).ReleasePatch()
	assert.Equal(t, "1.2.4", v.String())
}

func TestComplex_MarshalJSON(t *testing.T) {
	v, _ := newComplexFromString("1.2.3")

	type toJson struct {
		Complex Version
	}

	raw := toJson{
		Complex: v,
	}
	s, err := json.Marshal(raw)
	if err != nil {
		fmt.Println(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, "{\"Complex\":\"1.2.3\"}", string(s))
}
