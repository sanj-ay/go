package isogram

import (
	s "strings"
)

func IsIsogram(word string) bool {

	word = s.ToLower(word)
	m := make(map[rune]bool)

	for _, char := range word {

		if char == '-' || char == ' ' {
			continue
		}
		if _, ok := m[char];ok {
			return false
		}
		m[char] = true
	}
	return true
}
