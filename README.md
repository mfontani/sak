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

## Subcommands

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
