package main

import (
	"bufio"
	"regexp"
)

var rxANSI = regexp.MustCompile(`(\x1b\[[^m]*?m)`)

// StripANSI strips ANSI colors from the input file, and places the output in the output file.
func StripANSI(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		f := rxANSI.ReplaceAllString(l, "")
		w.Write([]byte(f))
	}
}
