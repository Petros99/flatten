package flatten

import (
	"errors"
	"strconv"
)

// Flatten generates a flat map from a nested one.  The original may include values of type map, slice and scalar,
// but not struct.  Keys in the flat map will be a compound of descending map keys and slice iterations.
// The separator between keys is set by separator.
func Flatten(nested map[string]interface{}, separator string) (map[string]interface{}, error) {
	flatmap := make(map[string]interface{}) // make a new map to hold the result
	return flatmap, flatten(true, flatmap, nested, "", separator)
}

var NotValidInputError = errors.New("not a valid input: map or slice")

// flatten takes a map and returns a new one where nested maps are replaced
// by separator keys.

func flatten(top bool, flatMap map[string]interface{}, nested interface{}, prefix string, separator string) error {
	assign := func(newKey string, v interface{}) error {
		switch v.(type) {
		case map[string]interface{}, []interface{}:
			if err := flatten(false, flatMap, v, newKey, separator); err != nil {
				return err
			}
		default:
			flatMap[newKey] = v
		}

		return nil
	}

	switch nested.(type) {
	case map[string]interface{}:
		for k, v := range nested.(map[string]interface{}) {
			newKey := enkey(top, prefix, k, separator)
			if err := assign(newKey, v); err != nil {
				return err
			}
		}
	case []interface{}:
		for i, v := range nested.([]interface{}) {
			newKey := enkey(top, prefix, strconv.Itoa(i), separator)
			if err := assign(newKey, v); err != nil {
				return err
			}
		}
	default:
		return NotValidInputError
	}

	return nil
}

func enkey(top bool, prefix, subkey string, separator string) string {
	key := prefix

	if top {
		key += subkey
	} else {
		key += separator + subkey
	}

	return key
}
