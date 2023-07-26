package main

import (
	"fmt"
	"os"
)

// Runes lists, one per line, the runes the given string is made up of
func Runes(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Need something to display.\n")
		os.Exit(1)
	}
	for _, arg := range args {
		shownHeader := false
		for _, r := range arg {
			if !shownHeader {
				fmt.Printf("Runes for %s:\n", arg)
				shownHeader = true
			}
			fmt.Printf("\t%s\n", describeRune(r))
		}
	}
}
