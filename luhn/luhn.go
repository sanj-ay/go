package luhn

import (
	"strings"
	"unicode"
)

func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	runes := []rune(number)
	sum := 0
	alternative := false
	if len(runes) < 2 {
		return false
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}
		d := int(runes[i] - '0')
		if alternative {
			if d *= 2; d > 9 {
				d -= 9
			}
		}
		sum += d
		alternative = !alternative
	}
	return sum%10 == 0
}
