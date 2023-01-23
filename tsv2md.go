package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// TSVToMD handles "csv2tsv [INPUT] [OUTPUT]"
func TSVToMD(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read header from TSV: %s\n", err)
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
			fmt.Fprintf(os.Stderr, "Failed to parse TSV: %s\n", err)
			os.Exit(1)
		}
		io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(row, " | ")))
	}
}
