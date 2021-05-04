package grains

import "errors"

func Total() uint64 {
	return 0x1p64 -1
}
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Invalid Number Of Squares")
	}
	return 1 << (n - 1), nil
}
