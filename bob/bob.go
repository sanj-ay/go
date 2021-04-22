package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if remark == "" {

		return "Fine. Be that way!"

	}

	if isYell(remark) && remark[len(remark)-1:] != "?" {

		return "Whoa, chill out!"
	}

	if isYellQuestion(remark) {

		return "Calm down, I know what I'm doing!"
	}

	if remark[len(remark)-1:] == "?" {

		return "Sure."
	}

	return "Whatever."
}

var charRegex = regexp.MustCompile(`[a-zA-Z]+`)

func isYell(remark string) bool {

	hasCharacters := charRegex.MatchString(remark)
	allUpper := strings.ToUpper(remark) == remark

	return hasCharacters && allUpper
}

func isYellQuestion(remark string) bool {

	hasCharacters := charRegex.MatchString(remark)

	allUpper := strings.ToUpper(remark) == remark

	questionMark := remark[len(remark)-1:] == "?"

	return hasCharacters && allUpper && questionMark
}
