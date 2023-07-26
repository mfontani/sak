package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
	"csv2tsv": {CSVToTSV, "Converts a CSV into a TSV", "csv2tsv [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a CSV file into a TSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"csv2md": {CSVToMD, "Converts a CSV to MarkDown", "csv2md [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a CSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"tsv2csv": {TSVToCSV, "Converts a TSV into a CSV", "tsv2csv [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a TSV file into a CSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"tsv2md": {TSVToMD, "Converts a TSV to MarkDown", "tsv2md [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a TSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.`},
	"xsv2md": {XSVToMD, "Converts a xSV to MarkDown", "xsv2md SEPARATOR [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Converts a xSV file into a MarkDown file. Requires a SEPARATOR to be given.
The SEPARATOR should be one-character, like $'\t' or ','.
Defaults to getting input from STDIN and giving output to STDOUT.
You can specify "-" for either INPUT_FILE or OUTPUT_FILE to mean STDIN and STDOUT,
respectively.
Accepts no options other than --help.`},
	"args": {ShowArgs, "Shows arguments given", "args [ARGUMENTS]",
		`Prints the number of arguments on STDERR, followed by a possibly colored "dump"
of each given argument on STDOUT. It highlights escape, backslash, space, tab,
newline and return carriage.
If an argument contains non-basic runes (outside of 0x20-0x7e, 0x09, 0x0a, 0x0d
and 0x1b, which are highlighted), it spits out one line per rune that comprises
the full string argument, and "describes" them, showing its decimal and hex
values, its name, and its properties.
Accepts no options other than --help.`},
	"rune": {Rune, "Shows runes matching the arguments", "rune [OPTIONS] ARGUMENT [ARGUMENT+]",
		`Prints/describes the Unicode runes matching ARGUMENT. Optionally shows them.
Uses the "fixed" rune descriptions for control characters, and support font
awesome runes as well.
ARGUMENT can be one of:
- a string, or a string starting with "+", is used to restrict runes to the ones
  which contain ARGUMENT in their description, case insensitively
- a string starting with "-", excludes runes whose description matches, case
  insensitively, ARGUMENT
- a string starting with "/" or with "+/", the ARGUMENT is taken as a case
  insensitive regular expression and runes are output if they match it
- a string starting with "-/", the ARGUMENT is taken as a case insensitive
  regular expression, and runes are excluded if they match it
- If only one ARGUMENT is given, and it matches a decimal number, then it
  displays information about the rune identified by that decimal number.
- If only one ARGUMENT is given, and it matches a hexadecimal number (starting
  with the string "0x"), then it displays information about the rune identified
  by that hexadecimal number.
Options:
    --help    Shows this help page
    --show    Shows the rune character as well as its description`},
	"runes": {Runes, "Shows runes contained in the string arguments", "runes ARGUMENT [ARGUMENT+]",
		`Prints/describes the Unicode runes contained in each ARGUMENT.
Uses the "fixed" rune descriptions for control characters, and supports font
awesome runes as well.  Useful to identify presence of combining characters.
Accepts no options other than --help.`},
	"stripansi": {StripANSI, "strips ansi from input", "stripansi [INPUT_FILE|-] [OUTPUT_FILE|-]",
		`Strips ANSI strings (i.e. \x1b[...m) from INPUT_FILE, and writes to OUTPUT_FILE.
Defaults to getting input from STDIN and giving output to STDOUT.
You can specify "-" for either INPUT_FILE or OUTPUT_FILE to mean STDIN and STDOUT,
respectively.
Accepts no options other than --help.`},
	"since": {Since, "Shows days, months etc between dates", "since START_DATE [END_DATE|TODAY]",
		`Prints the amount of days, weeks, months years between START_DATE and END_DATE.
END_DATE defaults to today's date. Dates need to be given in YYYY-MM-DD format.
DWIMs if END_DATE is <= START_DATE.
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
	// ... and show them sorted.
	keys := make([]string, 0, len(dispatch))
	l := 0
	for k := range dispatch {
		if len(k) > l {
			l = len(k)
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	lineFmt := fmt.Sprintf("  %%-%ds - %%s\n", l)
	for _, k := range keys {
		v, ok := dispatch[k]
		if ok {
			fmt.Printf(lineFmt, k, v.ShortDescription)
		}
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
		if which != "sak" && which != "--help" && which != "-help" {
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
