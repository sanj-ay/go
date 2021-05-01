package etl

import "strings"

func Transform(oldData map[int][]string) map[string]int {

	newData := make(map[string]int)

	for value, words := range oldData {
		for _, word := range words {
			newData[strings.ToLower(word)] = value
		}
	}

	return newData
}
