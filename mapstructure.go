package version

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"fmt"
)

func StringToVersionHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		from reflect.Type,
		to reflect.Type,
		data interface{}) (interface{}, error) {
		if from.Kind() != reflect.String {
			return data, nil
		}

		if to == reflect.TypeOf((*Version)(nil)).Elem() {
			return FromString(data.(string))
		}

		if to.Implements(reflect.TypeOf((*Version)(nil)).Elem()) {
			v, err := FromString(data.(string))
			if err != nil {
				return nil, err
			}

			valueType := reflect.TypeOf(v)

			if !valueType.ConvertibleTo(to) {
				return nil, fmt.Errorf("%s is unconvertible to %s", valueType.String(), to.String())
			}

			return v, nil
		}

		// Convert it by parsing
		return data, nil
	}
}
