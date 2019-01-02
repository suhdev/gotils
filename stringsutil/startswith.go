package stringsutil

import "strings"

// StartsWith checks whether a string ends with a given prefix
// this function is case sensitive
func StartsWith(base, prefix string) bool {
	l1 := len(base)
	l2 := len(prefix)
	if l1 < l2 {
		return false
	}
	x := base[0:l2]
	return x == prefix
}

// StartsWithInsensitive checks whether a string ends with a given prefix
// this function is case insensitive
func StartsWithInsensitive(base, prefix string) bool {
	l1 := len(base)
	l2 := len(prefix)
	if l1 < l2 {
		return false
	}
	x := strings.ToLower(base[0:l2])
	return x == strings.ToLower(prefix)
}
