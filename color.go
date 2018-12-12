package jpp

import (
	"fmt"

	au "github.com/logrusorgru/aurora"
)

// ColorScheme is used for specifying the output color or SGR
// using https://github.com/fatih/color
type ColorScheme struct {
	Null      ColoredFormat
	Bool      ColoredFormat
	Number    ColoredFormat
	String    ColoredFormat
	FieldName ColoredFormat
}

var (
	// DefaultScheme is plain color scheme that doesn't specify
	// any colors or SGRs.
	// Note that the default color scheme for CLI tool is not this.
	DefaultScheme = &ColorScheme{
		Null:      NoColor,
		Bool:      NoColor,
		Number:    NoColor,
		String:    NoColor,
		FieldName: NoColor,
	}
)

// ColoredFormat formats according to a format specifier and returns the resulting string.
type ColoredFormat = func(string, ...interface{}) string

// NoColor formats according to a format specifier and returns the resulting string.
func NoColor(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// Black formats according to a format specifier and returns the resulting string colored by black.
func Black(format string, args ...interface{}) string {
	return au.Sprintf(au.Black(format), args...)
}

// Red formats according to a format specifier and returns the resulting string colored by red.
func Red(format string, args ...interface{}) string {
	return au.Sprintf(au.Red(format), args...)
}

// Green formats according to a format specifier and returns the resulting string colored by green.
func Green(format string, args ...interface{}) string {
	return au.Sprintf(au.Green(format), args...)
}

// Brown formats according to a format specifier and returns the resulting string colored by brown.
func Brown(format string, args ...interface{}) string {
	return au.Sprintf(au.Brown(format), args...)
}

// Blue formats according to a format specifier and returns the resulting string colored by blue.
func Blue(format string, args ...interface{}) string {
	return au.Sprintf(au.Blue(format), args...)
}

// Magenta formats according to a format specifier and returns the resulting string colored by magenta.
func Magenta(format string, args ...interface{}) string {
	return au.Sprintf(au.Magenta(format), args...)
}

// Cyan formats according to a format specifier and returns the resulting string colored by cyan.
func Cyan(format string, args ...interface{}) string {
	return au.Sprintf(au.Cyan(format), args...)
}

// Gray formats according to a format specifier and returns the resulting string colored by gray.
func Gray(format string, args ...interface{}) string {
	return au.Sprintf(au.Gray(format), args...)
}


// BoldBlack formats according to a format specifier and returns the resulting bold string colored by black.
func BoldBlack(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BlackFg|au.BoldFm), args...)
}

// BoldRed formats according to a format specifier and returns the resulting bold string colored by red.
func BoldRed(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.RedFg|au.BoldFm), args...)
}

// BoldGreen formats according to a format specifier and returns the resulting bold string colored by green.
func BoldGreen(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.GreenFg|au.BoldFm), args...)
}

// BoldBrown formats according to a format specifier and returns the resulting bold string colored by brown.
func BoldBrown(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BrownFg|au.BoldFm), args...)
}

// BoldBlue formats according to a format specifier and returns the resulting bold string colored by blue.
func BoldBlue(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BlueFg|au.BoldFm), args...)
}

// BoldMagenta formats according to a format specifier and returns the resulting bold string colored by magenta.
func BoldMagenta(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.MagentaFg|au.BoldFm), args...)
}

// BoldCyan formats according to a format specifier and returns the resulting bold string colored by cyan.
func BoldCyan(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.CyanFg|au.BoldFm), args...)
}

// BoldGray formats according to a format specifier and returns the resulting bold string colored by gray.
func BoldGray(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.GrayFg|au.BoldFm), args...)
}
