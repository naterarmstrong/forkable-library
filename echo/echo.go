package echo

import (
	"fmt"
	"strings"
)

// Echo echoes a word so that it appears to echo in the terminal.
func Echo(s string) {
	// Maintainer's Note: Echoing multiple words is not in scope for this library and will not be supported.
	words := strings.Split(s, " ")
	if len(words) > 1 {
		panic("Only one word is supported!")
	}
	lowerCaseWord := strings.ToLower(words[0])
	noPunctuationWord := strings.TrimRightFunc(lowerCaseWord, func(r rune) bool {
		return r == '!' || r == '.'
	})
	fmt.Printf("%s\n%s...\n%s...\n", words[0], noPunctuationWord, noPunctuationWord)
}
