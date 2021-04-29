// Package raindrops implements integer to raindrop string conversion
package raindrops

import (
	"strconv"
)

// Convert converts integer to raindrop string
func Convert(t int) string {
	result := ""
	if t%3 == 0 {
		result += "Pling"
	}
	if t%5 == 0 {
		result += "Plang"
	}
	if t%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return strconv.Itoa(t)
	}
	return result
}
