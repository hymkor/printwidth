package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-runewidth"
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
	for _, s := range os.Args[1:] {
		report(s)
	}
}
