package jpp

import (
	"github.com/fatih/color"
)

// ColorScheme is used for specifying the output color or SGR
// using https://github.com/fatih/color
type ColorScheme struct {
	Null      *color.Color
	Bool      *color.Color
	Number    *color.Color
	String    *color.Color
	FieldName *color.Color
}

var (
	// DefaultScheme is plain color scheme that doesn't specify
	// any colors or SGRs.
	// Note that the default color scheme for CLI tool is not this.
	DefaultScheme = &ColorScheme{
		Null:      color.New(),
		Bool:      color.New(),
		Number:    color.New(),
		String:    color.New(),
		FieldName: color.New(),
	}
)
