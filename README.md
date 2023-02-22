# sak - my Swiss Army Knife

This is a multi-tool in one fat binary. Should have no dependencies.

`sak` can execute sub-commands, with arguments, described below. Example:

```bash
printf 'a\tb\tc\n1\t2\t3\n' | sak tsv2csv
a,b,c
1,2,3
```

If the tool is linked to the name of a sub-command, it'll function as if that
sub-command were invoked from the main script, i.e.

```bash
ln -s sak tsv2csv
printf 'a\tb\tc\n1\t2\t3\n' | ./tsv2csv
a,b,c
1,2,3
```

Tools might get added over time as I get around to it.  Starting small for the
moment, slowly adding ones I had lying around on a webapp somewhere so I can
have them _also_ available at the command line on most systems I have/use.

## How to install

You can download a pre-packaged binary from the
[releases page](https://github.com/mfontani/sak/releases),
or use a tool such as [ubi](https://github.com/houseabsolute/ubi) to fetch it
and install it to your local `$HOME/bin` (assuming it's in your `$PATH`, it'll
Just Work®).

```bash
$ ubi --project mfontani/sak --in "$HOME/bin"
$ sak --version
sak version v0.0.9
```

## Subcommands

### `args` - Shows arguments given

Synopsis: `args [ARGUMENTS]`

Prints the number of arguments on STDERR, followed by a possibly colored "dump"
of each given argument on STDOUT. It highlights escape, backslash, space, tab,
newline and return carriage.
If an argument contains non-basic runes (outside of 0x20-0x7e, 0x09, 0x0a, 0x0d
and 0x1b, which are highlighted), it spits out one line per rune that comprises
the full string argument, and "describes" them, showing its decimal and hex
values, its name, and its properties.
Accepts no options other than --help.

### `csv2md` - Converts a CSV to MarkDown

Synopsis: `csv2md [INPUT_FILE|-] [OUTPUT_FILE|-]`

Converts a CSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.

### `csv2tsv` - Converts a CSV into a TSV

Synopsis: `csv2tsv [INPUT_FILE|-] [OUTPUT_FILE|-]`

Converts a CSV file into a TSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.

### `rune` - Shows runes matching the arguments

Synopsis: `rune [OPTIONS] ARGUMENT [ARGUMENT+]`

Prints/describes the Unicode runes matching ARGUMENT. Optionally shows them.
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
    --show    Shows the rune character as well as its description

### `since` - Shows days, months etc between dates

Synopsis: `since START_DATE [END_DATE|TODAY]`

Prints the amount of days, weeks, months years between START_DATE and END_DATE.
END_DATE defaults to today's date. Dates need to be given in YYYY-MM-DD format.
DWIMs if END_DATE is <= START_DATE.
Accepts no options other than --help.

### `stripansi` - strips ansi from input

Synopsis: `stripansi [INPUT_FILE|-] [OUTPUT_FILE|-]`

Strips ANSI strings (i.e. \x1b[...m) from INPUT_FILE, and writes to OUTPUT_FILE.
Defaults to getting input from STDIN and giving output to STDOUT.
You can specify "-" for either INPUT_FILE or OUTPUT_FILE to mean STDIN and STDOUT,
respectively.
Accepts no options other than --help.

### `tsv2csv` - Converts a TSV into a CSV

Synopsis: `tsv2csv [INPUT_FILE|-] [OUTPUT_FILE|-]`

Converts a TSV file into a CSV file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.

### `tsv2md` - Converts a TSV to MarkDown

Synopsis: `tsv2md [INPUT_FILE|-] [OUTPUT_FILE|-]`

Converts a TSV file into a MarkDown file. Defaults to getting input from STDIN
and giving output to STDOUT. You can specify "-" for either INPUT_FILE or
OUTPUT_FILE to mean STDIN and STDOUT, respectively.
Accepts no options other than --help.

### `xsv2md` - Converts a xSV to MarkDown

Synopsis: `xsv2md SEPARATOR [INPUT_FILE|-] [OUTPUT_FILE|-]`

Converts a xSV file into a MarkDown file. Requires a SEPARATOR to be given.
The SEPARATOR should be one-character, like $'\t' or ','.
Defaults to getting input from STDIN and giving output to STDOUT.
You can specify "-" for either INPUT_FILE or OUTPUT_FILE to mean STDIN and STDOUT,
respectively.
Accepts no options other than --help.

## LICENSE

Copyright 2023 Marco Fontani <MFONTANI@cpan.org>

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
