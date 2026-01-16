package main

import (
	"flag"
	"fmt"

	"github.com/mattn/go-runewidth"
)

var (
	flagCode = flag.Int("n", -1, "code point")
)

func report(s string) {
	fmt.Printf("`%s` (%#v):\n", s, s)
	fmt.Printf("  StringWidth: %d\n", runewidth.StringWidth(s))
	rw := 0
	for _, c := range s {
		rw += runewidth.RuneWidth(c)
	}
	fmt.Printf("  Sum of RuneWidth: %d\n", rw)
}

func main() {
	flag.Parse()

	if *flagCode >= 0 {
		report(fmt.Sprintf("%c", *flagCode))
	}
	for _, s := range flag.Args() {
		report(s)
	}
}
