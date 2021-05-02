package grains

import "errors"

func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		n, _ := Square(i)
		total += n
	}
	return total

}
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Invalid Number Of Squares")
	}
	return 1 << (n - 1), nil
}
