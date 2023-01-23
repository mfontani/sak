package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// XSVToMD handles "xsv2md SEPARATOR [INPUT] [OUTPUT]"
func XSVToMD(args []string) {
	MaxArgs(3, args)
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "xsv2md requires a SEPARATOR, like $'\\t' or ','. See --help.\n")
		os.Exit(1)
	}
	if len(args[0]) == 0 {
		fmt.Fprintf(os.Stderr, "xsv2md requires a non-empty SEPARATOR, like $'\\t' or ','. See --help.\n")
		os.Exit(1)
	}
	separator := []rune(args[0])
	args = args[1:]
	r, w := IOArgs(args)
	reader := csv.NewReader(r)
	reader.Comma = separator[0]
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read header from xSV: %s\n", err)
		os.Exit(1)
	}
	io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(header, " | ")))
	for range header[:len(header)-1] {
		io.WriteString(w, "--- | ")
	}
	io.WriteString(w, " ---\n")
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse xSV: %s\n", err)
			os.Exit(1)
		}
		io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(row, " | ")))
	}
}

// CSVToMD handles "csv2md [INPUT] [OUTPUT]"
func CSVToMD(args []string) {
	args = append([]string{","}, args...)
	XSVToMD(args)
}

// TSVToMD handles "tsv2md [INPUT] [OUTPUT]"
func TSVToMD(args []string) {
	args = append([]string{"\t"}, args...)
	XSVToMD(args)
}
