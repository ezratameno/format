package format

import (
	"fmt"
	"strings"
)

// formatString will format prometheus label.
func formatString(s string) string {

	s = strings.ToLower(s)
	var res string
	for _, v := range s {
		if v > 'z' || v < 'a' {
			res = fmt.Sprintf("%s_", res)
		} else {
			res = fmt.Sprintf("%s%c", res, v)
		}
	}

	return res
}
