package gojson

import "strings"

func empty(s string) bool {
	return len(s) == 0
}

func startsWithUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	char := s[0:1]
	return char == strings.ToUpper(char)
}

func contains(elems []string, search string) bool {
	if (elems == nil) || (len(elems) == 0) {
		return false
	}
	for _, elem := range elems {
		if elem == search {
			return true
		}
	}
	return false
}
