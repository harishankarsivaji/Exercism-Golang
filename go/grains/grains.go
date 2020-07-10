package grains

import "errors"

//Square of uint64 range
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Bad argument")
	}
	return 1 << (uint(n) - 1), nil
}

// Total value of 2^64
func Total() uint64 {
	return ^uint64(0)
}
