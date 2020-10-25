package flatten

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	cases := []struct {
		test      string
		want      map[string]interface{}
		separator string
	}{
		{
			`{
				"foo": {
					"jim":"bean"
				},
				"fee": "bar",
				"n1": {
					"alist": [
						"a",
						"b",
						"c",
						{
							"d": "other",
							"e": "another"
						}
					]
				},
				"number": 1.4567,
				"bool":   true
			}`,
			map[string]interface{}{
				"foo.jim":      "bean",
				"fee":          "bar",
				"n1.alist.0":   "a",
				"n1.alist.1":   "b",
				"n1.alist.2":   "c",
				"n1.alist.3.d": "other",
				"n1.alist.3.e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			".",
		},
		{
			`{
				"foo": {
					"jim":"bean"
				},
				"fee": "bar",
				"n1": {
					"alist": [
					"a",
					"b",
					"c",
					{
						"d": "other",
						"e": "another"
					}
					]
				}
			}`,
			map[string]interface{}{
				"foo-jim":      "bean",
				"fee":          "bar",
				"n1-alist-0":   "a",
				"n1-alist-1":   "b",
				"n1-alist-2":   "c",
				"n1-alist-3-d": "other",
				"n1-alist-3-e": "another",
			},
			"-",
		},
		{
			`{
				"foo": {
					"jim":"bean"
				},
				"fee": "bar",
				"n1": {
					"alist": [
						"a",
						"b",
						"c",
						{
							"d": "other",
							"e": "another"
						}
					]
				},
				"number": 1.4567,
				"bool":   true
			}`,
			map[string]interface{}{
				"foo/jim":      "bean",
				"fee":          "bar",
				"n1/alist/0":   "a",
				"n1/alist/1":   "b",
				"n1/alist/2":   "c",
				"n1/alist/3/d": "other",
				"n1/alist/3/e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			"/",
		},
		{
			`{
				"foo": {
					"jim":"bean"
				},
				"fee": "bar",
				"n1": {
					"alist": [
						"a",
						"b",
						"c",
						{
							"d": "other",
							"e": "another"
						}
					]
				},
				"number": 1.4567,
				"bool":   true
			}`,
			map[string]interface{}{
				"foo_jim":      "bean",
				"fee":          "bar",
				"n1_alist_0":   "a",
				"n1_alist_1":   "b",
				"n1_alist_2":   "c",
				"n1_alist_3_d": "other",
				"n1_alist_3_e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			"_",
		},
	}

	for i, test := range cases {
		var m interface{}
		err := json.Unmarshal([]byte(test.test), &m)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
			continue
		}
		got, _ := Flatten(m.(map[string]interface{}), test.separator)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v wanted: %v", i+1, got, test.want)
		}
	}
}
