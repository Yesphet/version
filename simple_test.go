package version

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	"encoding/json"
)

func TestSimple_String(t *testing.T) {
	s1 := Simple(1)
	assert.Equal(t, "1", s1.String())

	s2 := Simple(math.MaxInt32)
	assert.Equal(t, "2147483647", s2.String())

	s3 := Simple(math.MaxInt64)
	assert.Equal(t, "9223372036854775807", s3.String())
}

func TestSimple_MarshalJSON(t *testing.T) {
	s1 := Simple(1)
	assert.Equal(t, "1", s1.String())

	js, err := json.Marshal(&s1)
	assert.Nil(t, err)

	assert.Equal(t, "\"1\"", string(js))
}

func TestSimple_Release(t *testing.T) {
	s1 := Simple(1)
	assert.Equal(t, "1", s1.String())

	s1.Release()
	assert.Equal(t, "2", s1.String())
}
