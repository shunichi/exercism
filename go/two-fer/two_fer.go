// Package twofer generates two-fer message
package twofer

import "fmt"

// ShareWith generates two-fer message
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
