package format

import (
	"fmt"
	"strings"
)

// formatString will format prometheus label.
func formatString(s string) string {

	s = strings.ToLower(s)
	var res string
	for _, r := range s {

		if !isLetter(r) && !isDigit(r) {
			res = fmt.Sprintf("%s_", res)
		} else {
			res = fmt.Sprintf("%s%c", res, r)
		}
	}

	return res
}

// isLetter checks if it's a lower case letter.
func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z'
}

// isDigit checks if it's a digit.
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
