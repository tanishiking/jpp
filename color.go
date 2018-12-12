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

type ColoredFormat = func(string, ...interface{}) string

func NoColor(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
func Black(format string, args ...interface{}) string {
	return au.Sprintf(au.Black(format), args...)
}
func Red(format string, args ...interface{}) string {
	return au.Sprintf(au.Red(format), args...)
}
func Green(format string, args ...interface{}) string {
	return au.Sprintf(au.Green(format), args...)
}
func Brown(format string, args ...interface{}) string {
	return au.Sprintf(au.Brown(format), args...)
}
func Blue(format string, args ...interface{}) string {
	return au.Sprintf(au.Blue(format), args...)
}
func Magenta(format string, args ...interface{}) string {
	return au.Sprintf(au.Magenta(format), args...)
}
func Cyan(format string, args ...interface{}) string {
	return au.Sprintf(au.Cyan(format), args...)
}
func Gray(format string, args ...interface{}) string {
	return au.Sprintf(au.Gray(format), args...)
}

func BoldBlack(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BlackFg|au.BoldFm), args...)
}
func BoldRed(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.RedFg|au.BoldFm), args...)
}
func BoldGreen(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.GreenFg|au.BoldFm), args...)
}
func BoldBrown(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BrownFg|au.BoldFm), args...)
}
func BoldBlue(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.BlueFg|au.BoldFm), args...)
}
func BoldMagenta(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.MagentaFg|au.BoldFm), args...)
}
func BoldCyan(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.CyanFg|au.BoldFm), args...)
}
func BoldGray(format string, args ...interface{}) string {
	return au.Sprintf(au.Colorize(format, au.GrayFg|au.BoldFm), args...)
}
