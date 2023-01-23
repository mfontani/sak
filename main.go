package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Version contains the binary version. This is added at build time.
var Version = "uncommitted"

type subcommand struct {
	Func             func([]string)
	ShortDescription string
	Synopsis         string
	FullDescription  string
}

var dispatch = map[string]subcommand{
	"csv2tsv": subcommand{CSVToTSV, "Converts a CSV into a TSV", "csv2tsv [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a CSV file into a TSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"csv2md": subcommand{CSVToMD, "Converts a CSV to MarkDown", "csv2md [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a CSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"tsv2csv": subcommand{TSVToCSV, "Converts a TSV into a CSV", "tsv2csv [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a TSV file into a CSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"tsv2md": subcommand{TSVToMD, "Converts a TSV to MarkDown", "tsv2md [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a TSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
}

// ShowVersion shows the version of this tool.
func ShowVersion() {
	fmt.Printf("sak version %s\n", Version)
}

// Help shows the help page for this command.
func Help() {
	ShowVersion()
	fmt.Println("Commands available:")
	// DWIM format string width to neatly show all commands
	l := 0
	for k := range dispatch {
		if len(k) > l {
			l = len(k)
		}
	}
	lineFmt := fmt.Sprintf("  %%-%ds - %%s\n", l)
	for k := range dispatch {
		fmt.Printf(lineFmt, k, dispatch[k].ShortDescription)
	}
	os.Exit(0)
}

func main() {
	// Support "exploding" the tool, i.e. "ln -s sak csv2tsv" makes it behave
	// like the "csv2tsv" subcommand.
	which := filepath.Base(os.Args[0])
	args := os.Args[1:]
	if _, ok := dispatch[which]; !ok && len(args) > 0 {
		which = args[0]
		args = args[1:]
	}
	sc, ok := dispatch[which]
	if !ok {
		if which == "--version" || which == "-version" {
			ShowVersion()
			os.Exit(0)
		}
		if which != "--help" && which != "-help" {
			fmt.Printf("No such command '%s'.\n", which)
		}
		Help()
	}
	// DWIM --help, -help etc.
	for _, arg := range args {
		if arg == "--version" || arg == "-version" {
			ShowVersion()
			os.Exit(0)
		}
		if arg == "--help" || arg == "-help" {
			fmt.Printf("Synopsis: %s\n", sc.Synopsis)
			fmt.Println(sc.FullDescription)
			os.Exit(0)
		}
	}
	sc.Func(args)
}
