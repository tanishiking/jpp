package jpp

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	p "github.com/tanishiking/prettier"
	"github.com/tidwall/gjson"
)

var (
	indent   string
	width    int
	coloring *ColorScheme
)

// Pretty prettifies specified json string.
func Pretty(jsonStr string, i string, w int, colorScheme *ColorScheme) (string, error) {
	if !gjson.Valid(jsonStr) {
		return "", errors.New("parse error: Invalid json input")
	}
	json := gjson.Parse(jsonStr)

	indent = i
	width = w
	if colorScheme != nil {
		coloring = colorScheme
	} else {
		coloring = DefaultScheme
	}

	var builder bytes.Buffer
	prettyRec(&builder, 0, json)
	return builder.String(), nil
}

func prettyRec(b *bytes.Buffer, depth int, j gjson.Result) {
	switch j.Type {
	case gjson.Null:
		color := coloring.Null
		b.WriteString(color("null"))
	case gjson.False:
		color := coloring.Bool
		b.WriteString(color("false"))
	case gjson.Number:
		color := coloring.Number
		b.WriteString(color("%v", j.Num))
	case gjson.String:
		color := coloring.String
		b.WriteString(color("\"%v\"", j.Str))
	case gjson.True:
		color := coloring.Bool
		b.WriteString(color("true"))
	case gjson.JSON:
		if j.IsArray() {
			items := j.Array()
			if allElemsAreScalar(items) {
				// Try to fit the json array in a single line
				// if all items are scalar values.
				indentLength := len([]rune(indent))
				sep := p.Concat([]p.Doc{p.Text(","), p.LineOrSpace()})
				ds := make([]p.Doc, 0, len(items))
				for _, item := range items {
					ds = append(ds, toDoc(item))
				}
				doc := p.TightBracketBy(
					p.Text("["),
					p.Text("]"),
					p.Intercalate(sep, ds),
					uint(indentLength),
				)
				layout := strings.Replace(
					p.Pretty(width-indentLength*depth, doc),
					"\n",
					fmt.Sprintf("\n%v", strings.Repeat(indent, depth)),
					-1,
				)
				b.WriteString(layout)
			} else {
				b.WriteString("[")
				depthInBracket := depth + 1
				newline(b, indent, depthInBracket)
				for i, item := range items {
					prettyRec(b, depthInBracket, item)
					if i != len(items)-1 {
						b.WriteString(",")
						newline(b, indent, depthInBracket)
					}
				}
				newline(b, indent, depth)
				b.WriteString("]")
			}
		} else {
			m := j.Map()
			if allValuesAreScalar(m) {
				// Try to fit the json object in a single line
				// if all values of the json object are scalar values.
				indentLength := len([]rune(indent))
				sep := p.Concat([]p.Doc{p.Text(","), p.LineOrSpace()})
				var kvs []p.Doc
				color := coloring.FieldName
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					length := len([]rune(k.Str))
					kv := p.Concat([]p.Doc{
						p.TextWithLength(color("\"%v\"", k.Str), length),
						p.Text(":"),
						p.Text(" "),
						toDoc(v),
					})
					kvs = append(kvs, kv)
					return true
				})
				doc := p.TightBracketBy(
					p.Text("{"),
					p.Text("}"),
					p.Fill(sep, kvs),
					uint(indentLength),
				)
				layout := strings.Replace(
					p.Pretty(width-indentLength*depth, doc),
					"\n",
					fmt.Sprintf("\n%v", strings.Repeat(indent, depth)),
					-1,
				)
				b.WriteString(layout)
			} else {
				b.WriteString("{")
				depthInBracket := depth + 1
				newline(b, indent, depthInBracket)
				len := len(m)
				i := 0
				color := coloring.FieldName
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					b.WriteString(color("\"%v\"", k.Str))
					b.WriteString(":")
					b.WriteString(" ")
					prettyRec(b, depthInBracket, v)
					if i != len-1 {
						b.WriteString(",")
						newline(b, indent, depthInBracket)
					}
					i++
					return true
				})
				newline(b, indent, depth)
				b.WriteString("}")
			}
		}
	}
}

func newline(dst *bytes.Buffer, indent string, depth int) {
	dst.WriteByte('\n')
	for i := 0; i < depth; i++ {
		dst.WriteString(indent)
	}
}

// toDoc convert gjson.Result to p.Doc
// note that we need to confirm that j is not gjson.JSON.
func toDoc(j gjson.Result) p.Doc {
	switch j.Type {
	default:
		return p.Empty()
	case gjson.Null:
		color := coloring.Null
		str := "null"
		length := len([]rune(str))
		return p.TextWithLength(color(str), length)
	case gjson.False:
		color := coloring.Bool
		str := "false"
		length := len([]rune(str))
		return p.TextWithLength(color(str), length)
	case gjson.Number:
		color := coloring.Number
		str := fmt.Sprintf("%v", j.Num)
		length := len([]rune(str))
		return p.TextWithLength(color(str), length)
	case gjson.String:
		color := coloring.String
		str := fmt.Sprintf("\"%v\"", j.Str)
		length := len([]rune(str))
		return p.TextWithLength(color(str), length)
	case gjson.True:
		color := coloring.Bool
		str := "true"
		length := len([]rune(str))
		return p.TextWithLength(color(str), length)
	}
}

func allElemsAreScalar(arr []gjson.Result) bool {
	for _, v := range arr {
		if v.Type == gjson.JSON {
			return false
		}
	}
	return true
}

func allValuesAreScalar(m map[string]gjson.Result) bool {
	for _, v := range m {
		if v.Type == gjson.JSON {
			return false
		}
	}
	return true
}
