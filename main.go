package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	"%a", "Mon",
	"%A", "Monday",
	"%b", "Jan",
	"%B", "January",
	"%c", time.RFC3339,
	"%C", "06",
	"%d", "02",
	"%C", "01/02/06",
	"%e", "_1/_2/_6",
	// "%E", "",
	"%F", "06-01-02",
	// "%G", "",
	// "%g", "",
	"%h", "Jan",
	"%H", "15",
	"%I", "03",
	// "%j", "",
	"%k", "3",
	"%l", "_3",
	"%m", "01",
	"%M", "04",
	"%n", "\n",
	// "%O", "",
	"%p", "PM",
	"%P", "pm",
	"%r", "03:04:05 PM",
	"%R", "03:04",
	// "%s", "",
	"%S", "05",
	"%t", "\t",
	"%T", "15:04:05",
	// "%u", "",
	// "%U", "",
	// "%V", "",
	// "%W", "",
	// "%x", "",
	// "%X", "",
	"%y", "06",
	"%Y", "2006",
	"%z", "-0700",
	"%Z", "MST",
	// "%+", "",
	"%%", "%",
)

func ts(ts string, scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Printf("%s%s\n", time.Now().Format(ts), scanner.Text())
	}
}

func main() {
	format := time.Stamp + " "
	if len(os.Args) > 1 {
		format = replacer.Replace(os.Args[1])
	}
	if len(os.Args) > 2 {
		for _, arg := range os.Args[2:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s: %s\n", os.Args[0], arg, err)
				os.Exit(1)
			}
			ts(format, bufio.NewScanner(f))
			f.Close()
		}
	} else {
		ts(format, bufio.NewScanner(os.Stdin))
	}
}
