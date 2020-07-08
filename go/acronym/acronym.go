package acronym

import (
	"strings"
)

// Abbreviate - to convert long name to acronym
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	result := ""
	for _, item := range strings.Fields(s) {
		first := string([]rune(item)[0])
		result += strings.ToUpper(first)
	}
	return result
}
