package jpp_test

import (
	"fmt"
	"testing"

	"github.com/tanishiking/jpp"
)

func Example() {
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

func TestPretty_PreserveOrder(t *testing.T) {
	jsonStr := `
{
  "foo": 1,
  "bar": 2,
  "baz": 3,
  "hello": 4,
  "world": 5,
  "numbers": [1,2,3,4,5]
}`
	actual, _ := jpp.Pretty(jsonStr, "  ", 20, nil)
	expected := `{
  "foo": 1,
  "bar": 2,
  "baz": 3,
  "hello": 4,
  "world": 5,
  "numbers": [
    1, 2, 3, 4, 5
  ]
}`
	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestPretty_Indent(t *testing.T) {
	jsonStr := `
{
  "characters": [
    {
      "name": "foo",
      "age": 25,
      "description": "bar"
    },
    {
      "name": "baz",
      "age": 100,
      "description": "foo"
    }
  ],
  "title": "foobar",
  "flag1": true,
  "flag2": false
}`
	actual, _ := jpp.Pretty(jsonStr, "    ", 20, nil)
	expected := `{
    "characters": [
        {
            "name": "foo",
            "age": 25,
            "description": "bar"
        },
        {
            "name": "baz",
            "age": 100,
            "description": "foo"
        }
    ],
    "title": "foobar",
    "flag1": true,
    "flag2": false
}`
	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestPretty_Width(t *testing.T) {
	jsonStr := `
{
  "characters": [
    {
      "name": "foo",
      "age": 25,
      "description": "bar"
    },
    {
      "name": "baz",
      "age": 100,
      "description": "foo"
    }
  ],
  "title": "foobar",
  "flag1": true,
  "flag2": false
}`
	actual, _ := jpp.Pretty(jsonStr, "    ", 100, nil)
	expected := `{
    "characters": [
        {
            "name": "foo", "age": 25, "description": "bar"
        },
        {
            "name": "baz", "age": 100, "description": "foo"
        }
    ],
    "title": "foobar",
    "flag1": true,
    "flag2": false
}`
	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
