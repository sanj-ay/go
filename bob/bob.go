package bob

import (
	"strings"
)

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	if remark == "" {

		return "Fine. Be that way!"

	}

	if isYell(remark) && Question(remark) {

		return "Calm down, I know what I'm doing!"
	}

	if isYell(remark) {

		return "Whoa, chill out!"
	}

	if Question(remark) {

		return "Sure."
	}

	return "Whatever."
}

func isYell(remark string) bool {

	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}

func Question(remark string) bool {

	return remark[len(remark)-1:] == "?"

}
