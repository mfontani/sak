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
