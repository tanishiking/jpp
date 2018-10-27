package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_width20(t *testing.T) {
	input := `[
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 ],
  [ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20 ]
]`
	inStream := strings.NewReader(input)
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	args := strings.Split("jpp -w 20", " ")
	c := &cli{
		inStream:  inStream,
		outStream: outStream,
		errStream: errStream,
	}

	c.run(args)

	expected := `[
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
`

	if outStream.String() != expected {
		t.Errorf("actual=%v, expected: %v", outStream.String(), expected)
	}
}
