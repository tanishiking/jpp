package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	au "github.com/logrusorgru/aurora"
	"github.com/tanishiking/jpp"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	defaultNull      = jpp.NoColor
	defaultBool      = jpp.NoColor
	defaultNumber    = jpp.NoColor
	defaultString    = jpp.Green
	defaultFieldName = jpp.BoldBlue
)

var (
	defaultCLIScheme = &jpp.ColorScheme{
		Null:      getColor("JPP_NULL", defaultNull),
		Bool:      getColor("JPP_BOOL", defaultBool),
		Number:    getColor("JPP_NUMBER", defaultNumber),
		String:    getColor("JPP_STRING", defaultString),
		FieldName: getColor("JPP_FIELDNAME", defaultFieldName),
	}
	monochrome = &jpp.ColorScheme{
		Null:      jpp.NoColor,
		Bool:      jpp.NoColor,
		Number:    jpp.NoColor,
		String:    jpp.NoColor,
		FieldName: jpp.NoColor,
	}
)

func getColor(envvar string, fallback jpp.ColoredFormat) jpp.ColoredFormat {
	v := os.Getenv(envvar)
	if v != "" {
		switch v {
		case "black":
			return jpp.Black
		case "red":
			return jpp.Red
		case "green":
			return jpp.Green
		case "brown":
			return jpp.Brown
		case "blue":
			return jpp.Blue
		case "magenta":
			return jpp.Magenta
		case "cyan":
			return jpp.Cyan
		case "gray":
			return jpp.Gray
		case "bold_black":
			return jpp.BoldBlack
		case "bold_red":
			return jpp.BoldRed
		case "bold_green":
			return jpp.BoldGreen
		case "bold_brown":
			return jpp.BoldBrown
		case "bold_blue":
			return jpp.BoldBlue
		case "bold_magenta":
			return jpp.BoldMagenta
		case "bold_cyan":
			return jpp.BoldCyan
		case "bold_gray":
			return jpp.BoldGray
		default:
			i, err := strconv.Atoi(v)
			if err != nil {
				return fallback
			}
			c := au.Color(i)
			if !c.IsValid() {
				return fallback
			}
			f := func(format string, args ...interface{}) string {
				return au.Sprintf(au.Colorize(format, c), args...)
			}
			return f
		}
	}
	return fallback
}

type cli struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

func (c *cli) run(args []string) int {
	var termErr error
	termWidth := -1
	f, ok := c.outStream.(*os.File)
	if ok {
		fd := int(f.Fd())
		terminalWidth, _, termErr := terminal.GetSize(fd)
		if termErr == nil {
			termWidth = terminalWidth
		}
	}

	var (
		indent  string
		width   int
		noColor bool
	)

	flags := flag.NewFlagSet("jpp", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.IntVar(&width, "w", termWidth, "width")
	flags.StringVar(&indent, "i", "  ", "indentation")
	flags.BoolVar(&noColor, "no-color", false, "disable the output color")
	err := flags.Parse(args[1:])
	if err != nil {
		return 1
	}

	if termErr != nil && width < 0 {
		fmt.Fprintln(c.errStream, "Couldn't read terminal width from your terminal.")
		return 1
	}

	reader := bufio.NewReader(c.inStream)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	jsonStr := string(output)

	var colorScheme *jpp.ColorScheme
	if noColor {
		colorScheme = monochrome
	} else {
		colorScheme = defaultCLIScheme
	}

	res, err := jpp.Pretty(jsonStr, indent, width, colorScheme)
	if err != nil {
		fmt.Fprintln(c.errStream, err.Error())
		return 1
	}
	fmt.Fprintln(c.outStream, res)
	return 0
}
