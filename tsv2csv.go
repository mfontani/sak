package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// TSVToCSV handles "tsv2csv [INPUT] [OUTPUT]"
func TSVToCSV(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	writer := csv.NewWriter(w)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse TSV: %s\n", err)
			os.Exit(1)
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write CSV: %s\n", err)
			os.Exit(1)
		}
	}
	writer.Flush()
}
