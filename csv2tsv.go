package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// CSVToTSV handles "csv2tsv [INPUT] [OUTPUT]"
func CSVToTSV(args []string) {
	MaxArgs(2, args)
	r, w := IOArgs(args)
	reader := csv.NewReader(r)
	writer := csv.NewWriter(w)
	writer.Comma = '\t'
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse CSV: %s\n", err)
			os.Exit(1)
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write TSV: %s\n", err)
			os.Exit(1)
		}
	}
	writer.Flush()
}
