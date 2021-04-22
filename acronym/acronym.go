package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {

	r, _ := regexp.Compile("([a-z'A-Z]+)")
	matches := r.FindAllString(s, -1)
	acronym := ""
	for i := 0; i < len(matches); i++ {
		acronym = acronym + strings.ToUpper(string(matches[i][0]))
	}
	return acronym
}
