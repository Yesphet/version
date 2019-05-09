package version

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEasyUpgrade(t *testing.T) {

	simple, _ := newSimpleFromString("1")
	EasyUpgrade(simple)
	assert.Equal(t, "2", simple.String())

	complex, _ := newComplexFromString("1.1.1")
	EasyUpgrade(complex)
	assert.Equal(t, "1.1.2", complex.String())

}
