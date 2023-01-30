package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

var rxANSI = regexp.MustCompile(`(\x1b\[[^m]*?m)`)

// StripANSI strips ANSI colors from the input file, and places the output in the output file.
func StripANSI(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	s := bufio.NewReader(r)
	for {
		l, err := s.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintf(os.Stderr, "Failed to read: %s\n", err)
			os.Exit(1)
		}
		f := rxANSI.ReplaceAllString(l, "")
		w.Write([]byte(f))
	}
}
