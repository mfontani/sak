package main

import (
	"fmt"
	"io"
	"os"
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

func argQuote(arg string, wantsColors bool) string {
	argRunes := []rune(arg)
	newRunes := make([]rune, 0)
	replacementRunes := replacementRunesColored
	if !wantsColors {
		replacementRunes = replacementRunesPlain
	}
	for _, argRune := range argRunes {
		rs, ok := replacementRunes[argRune]
		if ok {
			for _, r := range rs {
				newRunes = append(newRunes, r)
			}
		} else {
			newRunes = append(newRunes, argRune)
		}
	}
	return string(newRunes)
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
		fmt.Printf(fmtOutput, i+1, argQuote(v, wantsColors))
	}
}
