package raindrops

import (
	"strconv"
)

func Convert(a int) string {
	result := ""
	if a%3 == 0 {
		result = result + "Pling"

	}
	if a%5 == 0 {
		result = result + "Plang"
	}
	if a%7 == 0 {
		result = result + "Plong"
	}
	if result == "" {
		result = "" + strconv.Itoa(a)
	}

	return result
}
