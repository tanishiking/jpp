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
	indent string
	width  int
)

// Pretty prettifies specified json string.
func Pretty(jsonStr string, i string, w int) (string, error) {
	if !gjson.Valid(jsonStr) {
		return "", errors.New("parse error: Invalid json input")
	}
	json := gjson.Parse(jsonStr)

	indent = i
	width = w

	var builder bytes.Buffer
	prettyRec(&builder, 0, json)
	return builder.String(), nil
}

func prettyRec(b *bytes.Buffer, depth int, j gjson.Result) {
	switch j.Type {
	case gjson.Null:
		b.WriteString("null")
	case gjson.False:
		b.WriteString("false")
	case gjson.Number:
		b.WriteString(fmt.Sprintf("%v", j.Num))
	case gjson.String:
		b.WriteString(fmt.Sprintf("\"%v\"", j.Str))
	case gjson.True:
		b.WriteString("true")
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
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					kv := p.Concat([]p.Doc{
						p.Text(fmt.Sprintf("\"%v\"", k.Str)),
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
				j.ForEach(func(k gjson.Result, v gjson.Result) bool {
					b.WriteString(fmt.Sprintf("\"%v\"", k.Str))
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
		return p.Text("null")
	case gjson.False:
		return p.Text("false")
	case gjson.Number:
		return p.Text(fmt.Sprintf("%v", j.Num))
	case gjson.String:
		return p.Text(fmt.Sprintf("\"%v\"", j.Str))
	case gjson.True:
		return p.Text("true")
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
