package bob

import (
	"strings"
)

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if remark == "" {

		return "Fine. Be that way!"

	} else if isYell(remark) && IsQuestion(remark) {

		return "Calm down, I know what I'm doing!"
	} else if isYell(remark) {

		return "Whoa, chill out!"
	} else if IsQuestion(remark) {

		return "Sure."
	}

	return "Whatever."
}

func isYell(remark string) bool {

	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}

func IsQuestion(remark string) bool {

	return remark[len(remark)-1:] == "?"

}
