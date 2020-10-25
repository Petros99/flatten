// Flatten makes flat, one-dimensional maps from arbitrarily nested ones.
//
// It turns map keys into compound
// names, using given separator
//
// It takes input as Go maps.
//
// For example:
//	nested := map[string]interface{}{
//		"a": "b",
//		"c": map[string]interface{}{
//			"d": "e",
//			"f": "g",
//		},
//		"z": 1.4567,
//	}
//
//	flat := flatten.Flatten(nested, ".")
//
//	// output:
//	// map[string]interface{}{
//	//	"a":    "b",
//	//	"c.d":  "e",
//	//	"c.f":  "g",
//	//	"z":    1.4567,
//	// }
//

package flatten
