package hamming

import (
	"errors"
)

// Distance compare the len of string and count
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("")
	}

	counter := 0
	for i := 0; i < len(a); i++ {
		// Slice reference
		if a[i:i+1] != b[i:i+1] {
			counter++
		}
	}

	return counter, nil
}
