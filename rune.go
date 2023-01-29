package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/unicode/rangetable"
	"golang.org/x/text/unicode/runenames"
)

// "fix" things by providing more proper names for well-known ones that
// "runenames.Name(r)" insists on calling "<control>".
var fixedRuneNames = map[rune]string{
	0:   "NULL",
	1:   "START OF HEADING",
	2:   "START OF TEXT",
	3:   "END OF TEXT",
	4:   "END OF TRANSMISSION",
	5:   "ENQUIRY",
	6:   "ACKNOWLEDGE",
	7:   "ALERT",
	8:   "BACKSPACE",
	9:   "CHARACTER TABULATION",
	10:  "LINE FEED",
	11:  "LINE TABULATION",
	12:  "FORM FEED",
	13:  "CARRIAGE RETURN",
	14:  "SHIFT OUT",
	15:  "SHIFT IN",
	16:  "DATA LINK ESCAPE",
	17:  "DEVICE CONTROL ONE",
	18:  "DEVICE CONTROL TWO",
	19:  "DEVICE CONTROL THREE",
	20:  "DEVICE CONTROL FOUR",
	21:  "NEGATIVE ACKNOWLEDGE",
	22:  "SYNCHRONOUS IDLE",
	23:  "END OF TRANSMISSION BLOCK",
	24:  "CANCEL",
	25:  "END OF MEDIUM",
	26:  "SUBSTITUTE",
	27:  "ESCAPE",
	28:  "INFORMATION SEPARATOR FOUR",
	29:  "INFORMATION SEPARATOR THREE",
	30:  "INFORMATION SEPARATOR TWO",
	31:  "INFORMATION SEPARATOR ONE",
	127: "DELETE",
	128: "PADDING CHARACTER",
	129: "HIGH OCTET PRESET",
	130: "BREAK PERMITTED HERE",
	131: "NO BREAK HERE",
	132: "INDEX",
	133: "NEXT LINE",
	134: "START OF SELECTED AREA",
	135: "END OF SELECTED AREA",
	136: "CHARACTER TABULATION SET",
	137: "CHARACTER TABULATION WITH JUSTIFICATION",
	138: "LINE TABULATION SET",
	139: "PARTIAL LINE FORWARD",
	140: "PARTIAL LINE BACKWARD",
	141: "REVERSE LINE FEED",
	142: "SINGLE SHIFT TWO",
	143: "SINGLE SHIFT THREE",
	144: "DEVICE CONTROL STRING",
	145: "PRIVATE USE ONE",
	146: "PRIVATE USE TWO",
	147: "SET TRANSMIT STATE",
	148: "CANCEL CHARACTER",
	149: "MESSAGE WAITING",
	150: "START OF GUARDED AREA",
	151: "END OF GUARDED AREA",
	152: "START OF STRING",
	153: "SINGLE GRAPHIC CHARACTER INTRODUCER",
	154: "SINGLE CHARACTER INTRODUCER",
	155: "CONTROL SEQUENCE INTRODUCER",
	156: "STRING TERMINATOR",
	157: "OPERATING SYSTEM COMMAND",
	158: "PRIVACY MESSAGE",
	159: "APPLICATION PROGRAM COMMAND",
}

// Rune shows runes matching the arguments
func Rune(args []string) {
	showRune := false
	if len(args) > 0 {
		if args[0] == "--show" {
			showRune = true
			args = args[1:]
		}
	}
	needsContain := []string{}
	shouldntContain := []string{}
	needsMatch := []*regexp.Regexp{}
	shouldntMatch := []*regexp.Regexp{}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Need something to search for.\n")
		os.Exit(1)
	}
	for _, arg := range args {
		if strings.HasPrefix(arg, "+/") || strings.HasPrefix(arg, "/") {
			origArg := arg
			if strings.HasPrefix(arg, "+/") {
				arg = strings.TrimPrefix(arg, "+/")
			} else {
				arg = strings.TrimPrefix(arg, "/")
			}
			if strings.HasSuffix(arg, "/") {
				arg = strings.TrimSuffix(arg, "/")
			}
			if len(arg) == 0 {
				fmt.Fprintf(os.Stderr, "Too short regexp from %s\n", origArg)
				os.Exit(1)
			}
			rx, err := regexp.Compile("(?i)" + arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not compile regexp from %s: %s\n", origArg, err)
				os.Exit(1)
			}
			needsMatch = append(needsMatch, rx)
		} else if strings.HasPrefix(arg, "-/") {
			origArg := arg
			arg = strings.TrimPrefix(arg, "-/")
			if strings.HasSuffix(arg, "/") {
				arg = strings.TrimSuffix(arg, "/")
			}
			if len(arg) == 0 {
				fmt.Fprintf(os.Stderr, "Too short regexp from %s\n", origArg)
				os.Exit(1)
			}
			rx, err := regexp.Compile("(?i)" + arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not compile regexp from %s: %s\n", origArg, err)
				os.Exit(1)
			}
			shouldntMatch = append(shouldntMatch, rx)
		} else if strings.HasPrefix(arg, "-") {
			origArg := arg
			arg = strings.TrimPrefix(arg, "-")
			if len(arg) == 0 {
				fmt.Fprintf(os.Stderr, "Too short argument from %s\n", origArg)
				os.Exit(1)
			}
			shouldntContain = append(shouldntContain, strings.ToLower(arg))
		} else {
			origArg := arg
			if strings.HasPrefix(arg, "+") {
				arg = strings.TrimPrefix(arg, "+")
			}
			if len(arg) == 0 {
				fmt.Fprintf(os.Stderr, "Too short argument from %s\n", origArg)
				os.Exit(1)
			}
			needsContain = append(needsContain, strings.ToLower(arg))
		}
	}
	rt := rangetable.Assigned(runenames.UnicodeVersion)
	rangetable.Visit(rt, func(r rune) {
		name := runenames.Name(r)
		_, ok := fixedRuneNames[r]
		if ok {
			name = fixedRuneNames[r]
		}
		_, ok = faRunes[r]
		if ok {
			versionString := ""
			if faRunes[r].FromVersion != "" {
				versionString = fmt.Sprintf(" from version %s", faRunes[r].FromVersion)
			}
			name = fmt.Sprintf("%s - %s%s", name, faRunes[r].Name, versionString)
		}
		_, ok = nfRunes[r]
		if ok {
			name = fmt.Sprintf("%s - %s", name, nfRunes[r])
		}
		name = strings.ToLower(name)
		for i := 0; i < len(needsContain); i++ {
			if !strings.Contains(name, needsContain[i]) {
				return
			}
		}
		for i := 0; i < len(needsMatch); i++ {
			if !needsMatch[i].MatchString(name) {
				return
			}
		}
		for i := 0; i < len(shouldntContain); i++ {
			if strings.Contains(name, shouldntContain[i]) {
				return
			}
		}
		for i := 0; i < len(shouldntMatch); i++ {
			if shouldntMatch[i].MatchString(name) {
				return
			}
		}
		if showRune {
			fmt.Printf("%s - %s\n", string(r), describeRune(r))
		} else {
			fmt.Println(describeRune(r))
		}
	})
}
