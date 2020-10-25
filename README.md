flatten
=======

[![GoDoc](https://godoc.org/github.com/peterupton/flatten?status.png)](https://godoc.org/github.com/peterupton/flatten)
[![Build Status](https://travis-ci.org/peterupton/flatten.svg?branch=master)](https://travis-ci.org/peterupton/flatten)

Flatten makes flat, one-dimensional maps from arbitrarily nested ones.

It turns map keys into compound names.

It takes input as Go structures.



```go
nested := map[string]interface{}{
   "a": "b",
   "c": map[string]interface{}{
       "d": "e",
       "f": "g",
   },
   "z": 1.4567,
}

flat, err := flatten.Flatten(nested, ".")

// output:
// map[string]interface{}{
//  "a":    "b",
//  "c.d": "e",
//  "c.f": "g",
//  "z":    1.4567,
// }
```


See [godoc](https://godoc.org/github.com/peterupton/flatten) for API.
