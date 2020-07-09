package luhn

import "strings"

func Valid(str string) bool {
	str = strings.Replace(str, " ", "", -1)

	if len(str) <= 1 {
		return false
	}

	var digits []int

	// check non-digit characters
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
		digits = append(digits, int(r-'0'))
	}

	// double
	for i := len(digits) - 2; i >= 0; i -= 2 {
		n := digits[i] * 2
		if n > 9 {
			n = n - 9
		}
		digits[i] = n
	}

	sum := 0
	for _, n := range digits {
		sum += n
	}
	return sum%10 == 0
}
