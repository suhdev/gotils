package stringsutil

import "strings"

// EndsWith checks whether a string ends with a given postfix
// this function is case sensitive
func EndsWith(base, postfix string) bool {
	l1 := len(base)
	l2 := len(postfix)
	if l1 < l2 {
		return false
	}
	x := base[l1-l2:]
	return x == postfix
}

// EndsWithInsensitive checks whether a string ends with a given postfix
// this function is case insensitive
func EndsWithInsensitive(base, postfix string) bool {
	l1 := len(base)
	l2 := len(postfix)
	if l1 < l2 {
		return false
	}
	x := strings.ToLower(base[l1-l2:])
	return x == strings.ToLower(postfix)
}
