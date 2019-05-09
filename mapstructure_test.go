package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"reflect"
)

func TestStringToVersionHookFunc_simple_success(t *testing.T) {
	f := StringToVersionHookFunc()
	v, err := f(reflect.TypeOf("1"), reflect.TypeOf((*Simple)(nil)), "1")
	assert.Nil(t, err)

	assert.IsType(t, (*Simple)(nil), v)

	assert.Equal(t, "1", v.(*Simple).String())
}

func TestStringToVersionHookFunc_complex_success(t *testing.T) {
	f := StringToVersionHookFunc()
	v, err := f(reflect.TypeOf("1"), reflect.TypeOf(&Complex{}), "1.1.1")
	assert.Nil(t, err)

	assert.IsType(t, (*Complex)(nil), v)

	assert.Equal(t, "1.1.1", v.(*Complex).String())
}

func TestStringToVersionHookFunc_version_success(t *testing.T) {
	f := StringToVersionHookFunc()
	v, err := f(reflect.TypeOf("1"), reflect.TypeOf((*Version)(nil)).Elem(), "1.1.1")
	assert.Nil(t, err)

	assert.IsType(t, (*Complex)(nil), v)

	assert.Equal(t, "1.1.1", v.(Version).String())
}

func TestStringToVersionHookFunc_fail(t *testing.T) {
	f := StringToVersionHookFunc()
	v, err := f(reflect.TypeOf("1"), reflect.TypeOf((*Complex)(nil)), "1")
	assert.Nil(t, v)
	assert.Error(t, err)

}
