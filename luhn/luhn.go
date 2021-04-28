package luhn

import (
	"strings"
	"unicode"
)

func Valid(number string)bool{
	number = strings.ReplaceAll(number, " ", "")
	runes := []rune(number)
	sum := 0
	alternative := false
	if len(runes) <2{
		return false
	}
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]){
			return false
		}
		if alternative {
			d := runes[i] - '0'
			if d *=2; d> 9 {
				sum += int(d) - 9
			} else {
				sum += int(d)
			}
			alternative = !alternative
		} else {
			d := runes[i] - '0'
			sum = sum + int(d)
			alternative = !alternative
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}