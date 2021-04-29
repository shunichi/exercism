// Package hamming implements the Hamming Distance calculation
package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("different length")
	}
	distance := 0
	for i, c := range []byte(a) {
		if c != b[i] {
			distance++
		}
	}
	return distance, nil
}
