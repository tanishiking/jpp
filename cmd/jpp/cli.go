package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/tanishiking/jpp"
)

var (
	defaultNull      = color.New()
	defaultBool      = color.New()
	defaultNumber    = color.New()
	defaultString    = color.New(color.FgGreen)
	defaultFieldName = color.New(color.FgBlue, color.Bold)
)

var (
	defaultCLIScheme = &jpp.ColorScheme{
		Null:      getColor("JPP_NULL", defaultNull),
		Bool:      getColor("JPP_BOOL", defaultBool),
		Number:    getColor("JPP_NUMBER", defaultNumber),
		String:    getColor("JPP_STRING", defaultString),
		FieldName: getColor("JPP_FIELDNAME", defaultFieldName),
	}
)

func getColor(envvar string, fallback *color.Color) *color.Color {
	v := os.Getenv(envvar)
	if v != "" {
		i, err := strconv.Atoi(v)
		if err != nil {
			return fallback
		}
		attr := color.Attribute(i)
		return color.New(attr)
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
		indent string
		width  int
	)

	flags := flag.NewFlagSet("jpp", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.IntVar(&width, "w", termWidth, "width")
	flags.StringVar(&indent, "i", "  ", "indentation")
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
	res, err := jpp.Pretty(jsonStr, indent, width, defaultCLIScheme)
	if err != nil {
		fmt.Fprintln(c.errStream, err.Error())
		return 1
	}
	fmt.Fprintln(c.outStream, res)
	return 0
}
