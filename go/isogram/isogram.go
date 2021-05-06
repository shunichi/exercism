// Package isogram implements isogram checker
package isogram

import (
	"strings"
)

// IsIsogram check if stinng is isogram
func IsIsogram(s string) bool {
	letterMap := make(map[rune]bool)
	for _, c := range strings.ToLower(s) {
		if c == '-' || c == ' ' {
			continue
		}
		if letterMap[c] {
			return false
		}
		letterMap[c] = true
	}
	return true
}
