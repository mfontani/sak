package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// CSVToMD handles "csv2tsv [INPUT] [OUTPUT]"
func CSVToMD(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	reader := csv.NewReader(r)
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read header from CSV: %s\n", err)
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
			fmt.Fprintf(os.Stderr, "Failed to parse CSV: %s\n", err)
			os.Exit(1)
		}
		io.WriteString(w, fmt.Sprintf("%s\n", strings.Join(row, " | ")))
	}
}
