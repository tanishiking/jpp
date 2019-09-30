# jpp
[![Build Status](https://travis-ci.org/tanishiking/jpp.svg?branch=master)](https://travis-ci.org/tanishiking/jpp) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/b43969558e8246b4a2e5167eff17f21d)](https://app.codacy.com/app/tanishiking/jpp?utm_source=github.com&utm_medium=referral&utm_content=tanishiking/jpp&utm_campaign=Badge_Grade_Dashboard)

JSON Prettier Printer that occupies a minimal number of lines while pretty-printing given JSON, using [prettier](https://github.com/tanishiking/prettier) which is Go implementation of ["Wadler's "A Prettier Printer"](http://homepages.inf.ed.ac.uk/wadler/papers/prettier/prettier.pdf).

`jpp` is quite useful when we want to pretty print the JSON whose each node has a lot of children scalar values.

![screenshot from 2018-10-28 16-57-16](https://user-images.githubusercontent.com/9353584/47613438-bb96a700-dad2-11e8-872c-4309d4330aef.png)
This `example.json` cites from https://json.org/example.html

## Instalation
### Homebrew
```
$ brew install tanishiking/jpp/jpp
```

### Download binary from GitHub Releases
https://github.com/tanishiking/jpp/releases

### Build from source
```
$ go get -u github.com/tanishiking/jpp/cmd/jpp
```

### Develop
```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
$ dep ensure
$ go test .

## jpp command
### Options
- `-w`: width (default: your terminal width)
  - Note that this command does not guarantee there are no lines longer than `width`
  - It just attempts to keep lines within this length when possible.
- `-i`: indent string (default: `'  '`)

### Environment Variables
We can specify the color of output string using following environment variables

- `JPP_NULL`
- `JPP_BOOL`
- `JPP_NUMBER`
- `JPP_STRING`
- `JPP_FIELDNAME`

and builtin colors:

- `black`
- `red`
- `green`
- `brown`
- `blue`
- `magenta`
- `cyan`
- `gray`
- `bold_black`
- `bold_red`
- `bold_green`
- `bold_brown`
- `bold_blue`
- `bold_magenta`
- `bold_cyan`
- `bold_gray`

For example, `cat some.json | JPP_NUMBER=bold_red jpp`.
In addition, we can specify complicated styles using SGR code.
See: https://en.wikipedia.org/wiki/ANSI_escape_code

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
