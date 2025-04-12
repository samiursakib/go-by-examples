package main

import (
	"fmt"
	"unicode/utf8"
)

func stringsAndRunes() {
	const s = "สวัสดี"
	fmt.Println(s, "length:", len(s))
	for _, c := range s {
		fmt.Print(c, " ")
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString("é"), len("é"))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d", runeValue, idx)
		fmt.Println()
	}
	fmt.Println()

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d %d", runeValue, i, width)
		fmt.Println()
		w = width
	}
}
