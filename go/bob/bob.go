package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	convo := strings.TrimSpace(remark)

	switch {
	case convo == "":
		return "Fine. Be that way!"

	case isYelling(convo):
		return "Whoa, chill out!"

	case isYellingQuestion(convo):
		return "Calm down, I know what I'm doing!"

	case isQuestion(convo):
		return "Sure."

	default:
		return "Whatever."
	}
}

func isYellingQuestion(s string) bool {
	return isYelling(s) && isQuestion(s)
}

func isYelling(s string) bool {
	if removeNonLetters(s) == "" {
		return false
	}

	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	return strings.HasSuffix(s, "?")
}

func removeNonLetters(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}, s)
}
