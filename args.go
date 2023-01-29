package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

// MaxArgs checks that we have max `n` `args`, or dies.
func MaxArgs(n int, args []string) {
	if len(args) > n {
		fmt.Fprintf(os.Stderr, "Need max %d parameters, got %d.\n", n, len(args))
		os.Exit(1)
	}
}

// RequireArgs checks that we have `n` `args`, or dies.
func RequireArgs(n int, args []string) {
	if len(args) != n {
		fmt.Fprintf(os.Stderr, "Need %d parameters, got %d.\n", n, len(args))
		os.Exit(1)
	}
}

// IOArgs allows INPUT OUTPUT as parameters, or DWIMs STDIN/STDOUT
func IOArgs(args []string) (io.Reader, io.Writer) {
	r := os.Stdin
	w := os.Stdout
	if len(args) >= 1 {
		if args[0] != "-" {
			var err error
			r, err = os.Open(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not open %s for reading: %s\n", args[0], err)
				os.Exit(1)
			}
		}
	}
	if len(args) >= 2 {
		if args[1] != "-" {
			var err error
			w, err = os.Open(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not open %s for writing: %s\n", args[1], err)
				os.Exit(1)
			}
		}
	}
	return r, w
}

var replacementRunesColored = map[rune][]rune{
	'\\':   []rune("\x1b[31m\\\\\x1b[0m"),
	' ':    []rune("\x1b[30;46m\\ \x1b[0m"),
	'\t':   []rune("\x1b[30;44m\\t\x1b[0m"),
	'\n':   []rune("\x1b[30;45m\\n\x1b[0m"),
	'\r':   []rune("\x1b[30;45m\\r\x1b[0m"),
	'\x1b': []rune("\x1b[32;42m\\e\x1b[0m"),
}

var replacementRunesPlain = map[rune][]rune{
	'\\':   []rune("\\\\"),
	' ':    []rune("\\ "),
	'\t':   []rune("\\t"),
	'\n':   []rune("\\n"),
	'\r':   []rune("\\r"),
	'\x1b': []rune("\\e"),
}

// automatically turn colors on/off depending on the circumstances in which
// we're being called.
func getWantsColors() bool {
	wantsColors := false
	// If STDOUT is a terminal, we can have colors.
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		wantsColors = true
	}
	// If the user's specifically requesting no colors from any application, we
	// should honor their choice.
	if os.Getenv("NO_COLOR") != "" {
		wantsColors = false
	}
	// If the user's specifically requesting no colors from THIS application, we
	// should honor their choice, too.
	if os.Getenv("SAK_NO_COLOR") != "" {
		wantsColors = false
	}
	// A dumb terminal should never show colors.
	if os.Getenv("TERM") == "dumb" {
		wantsColors = false
	}
	return wantsColors
}

func isBasicRune(r rune) bool {
	// A "Basic" rune is:
	// - between 0x20 (space) and 0x7e (~)
	// - or "handled by replacementRunes", like:
	//   - 0x09, tab
	//   - 0x0a, LF
	//   - 0x0d, CR
	//   - 0x1b, Escape
	// Basic runes can either be printed out as-is, or have been properly
	// handled by replacementRunes.  Anything else should be output with
	// care, or we might want to "explode" it to show the user which runes
	// were actually output.
	if r >= 0x20 && r <= 0x7e {
		return true
	}
	_, ok := replacementRunesPlain[r]
	if ok {
		return true
	}
	return false
}

func argQuoteAndExplode(arg string, wantsColors bool) string {
	argRunes := []rune(arg)
	newRunes := make([]rune, 0)
	replacementRunes := replacementRunesColored
	if !wantsColors {
		replacementRunes = replacementRunesPlain
	}
	isAllBasic := true
	for _, argRune := range argRunes {
		rs, ok := replacementRunes[argRune]
		if ok {
			for _, r := range rs {
				newRunes = append(newRunes, r)
			}
		} else {
			if !isBasicRune(argRune) {
				isAllBasic = false
			}
			newRunes = append(newRunes, argRune)
		}
	}
	if isAllBasic {
		return string(newRunes)
	}
	// For args which don't "just" contain basic runes we'll return:
	// - string(newRunes) (the argument with known runes replaced)
	// - one indented line per rune, with "an explanation" for non-basic ones
	// ... joined by newlines.
	lines := make([]string, 0)
	lines = append(lines, string(newRunes))
	for _, r := range argRunes {
		lines = append(lines, describeRune(r))
	}
	return strings.Join(lines, "\n")
}

func describeRune(r rune) string {
	name := runenames.Name(r)
	_, ok := fixedRuneNames[r]
	if ok {
		name = fixedRuneNames[r]
	}
	_, ok = faRunes[r]
	if ok {
		versionString := ""
		if faRunes[r].FromVersion != "" {
			versionString = fmt.Sprintf(" from version %s", faRunes[r].FromVersion)
		}
		name = fmt.Sprintf("%s - %s%s", name, faRunes[r].Name, versionString)
	}
	_, ok = nfRunes[r]
	if ok {
		name = fmt.Sprintf("%s - %s", name, nfRunes[r])
	}
	var flags []string
	if unicode.IsControl(r) {
		flags = append(flags, "control")
	}
	if unicode.IsDigit(r) {
		flags = append(flags, "digit")
	}
	if unicode.IsGraphic(r) {
		flags = append(flags, "graphic")
	}
	if unicode.IsLetter(r) {
		flags = append(flags, "letter")
	}
	if unicode.IsLower(r) {
		flags = append(flags, "lower")
	}
	if unicode.IsMark(r) {
		flags = append(flags, "mark")
	}
	if unicode.IsNumber(r) {
		flags = append(flags, "number")
	}
	if unicode.IsPrint(r) {
		flags = append(flags, "printable")
	}
	if unicode.IsPunct(r) {
		flags = append(flags, "punct")
	}
	if unicode.IsSpace(r) {
		flags = append(flags, "space")
	}
	if unicode.IsSymbol(r) {
		flags = append(flags, "symbol")
	}
	if unicode.IsTitle(r) {
		flags = append(flags, "title")
	}
	if unicode.IsUpper(r) {
		flags = append(flags, "upper")
	}
	// decimal, hex, name, flags
	return fmt.Sprintf("  - %-4s 0x%04x %s (%s)", fmt.Sprintf("%d", r), r, name, strings.Join(flags, ", "))
}

// ShowArgs is the command that shows all arguments given to the program
func ShowArgs(args []string) {
	wantsColors := getWantsColors()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "%d ARGV.\n", len(args))
		return
	}
	fmt.Fprintf(os.Stderr, "%d ARGV:\n", len(args))
	fmtOutput := fmt.Sprintf("%%-%dd\t%%s\n", len(fmt.Sprintf("%d", len(args))))
	for i, v := range args {
		fmt.Printf(fmtOutput, i+1, argQuoteAndExplode(v, wantsColors))
	}
}
