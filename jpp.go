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
		b.WriteString(color.Sprint("null"))
	case gjson.False:
		color := coloring.Bool
		b.WriteString(color.Sprint("false"))
	case gjson.Number:
		color := coloring.Number
		b.WriteString(color.Sprintf("%v", j.Num))
	case gjson.String:
		color := coloring.String
		b.WriteString(color.Sprintf("\"%v\"", j.Str))
	case gjson.True:
		color := coloring.Bool
		b.WriteString(color.Sprint("true"))
	case gjson.JSON:
		if j.IsArray() {
			items := j.Array()
			b.WriteString("[")
			depthInBracket := depth + 1
			newline(b, indent, depthInBracket)
			if allElemsAreScalar(items) {
				indentLength := len([]rune(indent))
				sep := p.Concat([]p.Doc{p.Text(","), p.LineOrSpace()})
				ds := make([]p.Doc, 0, len(items))
				for _, item := range items {
					ds = append(ds, toDoc(item))
				}
				doc := p.Intercalate(sep, ds)
				layout := strings.Replace(
					p.Pretty(width-indentLength*depthInBracket, doc),
					"\n",
					fmt.Sprintf("\n%v", strings.Repeat(indent, depthInBracket)),
					-1,
				)
				b.WriteString(layout)
			} else {
				len := len(items)
				for i, item := range items {
					prettyRec(b, depthInBracket, item)
					if i != len-1 {
						b.WriteString(",")
						newline(b, indent, depthInBracket)
					}
				}
			}
			newline(b, indent, depth)
			b.WriteString("]")
		} else {
			b.WriteString("{")
			depthInBracket := depth + 1
			newline(b, indent, depthInBracket)
			m := j.Map()
			if allValuesAreScalar(m) {
				indentLength := len([]rune(indent))
				sep := p.Concat([]p.Doc{p.Text(","), p.LineOrSpace()})
				var kvs []p.Doc
				color := coloring.FieldName
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					length := len([]rune(k.Str))
					kv := p.Concat([]p.Doc{
						p.TextWithLength(color.Sprintf("\"%v\"", k.Str), length),
						p.Text(":"),
						p.Text(" "),
						toDoc(v),
					})
					kvs = append(kvs, kv)
					return true
				})
				doc := p.Fill(sep, kvs)
				layout := strings.Replace(
					p.Pretty(width-indentLength*depthInBracket, doc),
					"\n",
					fmt.Sprintf("\n%v", strings.Repeat(indent, depthInBracket)),
					-1,
				)
				b.WriteString(layout)
			} else {
				len := len(m)
				i := 0
				color := coloring.FieldName
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					b.WriteString(color.Sprintf("\"%v\"", k.Str))
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
			}
			newline(b, indent, depth)
			b.WriteString("}")
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
		return p.TextWithLength(color.Sprint(str), length)
	case gjson.False:
		color := coloring.Bool
		str := "false"
		length := len([]rune(str))
		return p.TextWithLength(color.Sprint(str), length)
	case gjson.Number:
		color := coloring.Number
		str := fmt.Sprintf("%v", j.Num)
		length := len([]rune(str))
		return p.TextWithLength(color.Sprint(str), length)
	case gjson.String:
		color := coloring.String
		str := fmt.Sprintf("\"%v\"", j.Str)
		length := len([]rune(str))
		return p.TextWithLength(color.Sprint(str), length)
	case gjson.True:
		color := coloring.Bool
		str := "true"
		length := len([]rune(str))
		return p.TextWithLength(color.Sprint(str), length)
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
