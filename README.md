# jpp [![Build Status](https://travis-ci.org/tanishiking/jpp.svg?branch=master)](https://travis-ci.org/tanishiking/jpp)
JSON Prettier Printer that occupies a minimal number of lines while pretty-printing given JSON, using [prettier](https://github.com/tanishiking/prettier) which is Go implementation of [Wadler's "A Prettier Printer"](http://homepages.inf.ed.ac.uk/wadler/papers/prettier/prettier.pdf).

## jpp command
### Options
- `-w`: width (default: your terminal width)
  - Note that this command does not guarantee there are no lines longer than `width`
  - It just attempts to keep lines within this length when possible.
- `-i`: indent string (default: `'  '`)

### Environment Variables
You can specify ANSI color to output colorized or SGR defined output to the standard output using following environment variables.

- `JPP_NULL`
- `JPP_BOOL`
- `JPP_NUMBER`
- `JPP_STRING`
- `JPP_FIELDNAME`

See: http://ascii-table.com/ansi-escape-sequences.php

```
$ go get -u github.com/tanishiking/jpp/cmd/jpp
$ cat numbers.json
[
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20 ]
]
$ cat numbers.json | jpp -w 20
[
  [
    1, 2, 3, 4, 5,
    6, 7, 8, 9, 10
  ],
  [
    1, 2, 3, 4, 5,
    6, 7, 8, 9, 10,
    11, 12, 13, 14,
    15
  ],
  [
    1, 2, 3, 4, 5,
    6, 7, 8, 9, 10,
    11, 12, 13, 14,
    15, 16, 17, 18,
    19, 20
  ]
]
```

## Package Usage
```go
import (
	"fmt"

	"github.com/tanishiking/jpp"
)

func main() {
	jsonStr := `
[
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20 ]
]
`
	res, _ := jpp.Pretty(jsonStr, "  ", 20, nil)
	fmt.Println(res)
	// [
	//   [
	//     1, 2, 3, 4, 5,
	//     6, 7, 8, 9, 10
	//   ],
	//   [
	//     1, 2, 3, 4, 5,
	//     6, 7, 8, 9, 10,
	//     11, 12, 13, 14,
	//     15
	//   ],
	//   [
	//     1, 2, 3, 4, 5,
	//     6, 7, 8, 9, 10,
	//     11, 12, 13, 14,
	//     15, 16, 17, 18,
	//     19, 20
	//   ]
	// ]
}
```
