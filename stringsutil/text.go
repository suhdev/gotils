package stringsutil

import (
	"regexp"
	"strings"
)

// Text a JS-like string type that implements
// the same methods of a JS string
type Text string

// ToLower converts the text to lower case
func (t *Text) ToLower() string {
	return strings.ToLower(string(*t))
}

// ToUpper converts the text to upper case
func (t *Text) ToUpper() string {
	return strings.ToUpper(string(*t))
}

// Trim trims white space around text
func (t *Text) Trim() string {
	return strings.TrimSpace(string(*t))
}

// Repeat repeats a string and return new string
func (t *Text) Repeat(count int) string {
	return strings.Repeat(string(*t), count)
}

// IndexOf returns the index of the first occurence of the
// provided text string
func (t *Text) IndexOf(text string) int {
	return strings.Index(string(*t), text)
}

// LastIndexOf returns the index of the last occurence of the
// provided text string
func (t *Text) LastIndexOf(text string) int {
	return strings.LastIndex(string(*t), text)
}

// ForEach loops through characters of string calling the provided
// callback on each character
func (t *Text) ForEach(cb func(string, int)) {
	for i, r := range *t {
		cb(string(r), i)
	}
}

// Replace replaces a given substring with a new one
func (t *Text) Replace(search, rep string) Text {
	return Text(strings.Replace(string(*t), rep, rep, -1))
}

// ReplaceRegex replaces string based on a regex
func (t *Text) ReplaceRegex(r *regexp.Regexp, rep string) Text {
	return Text(r.ReplaceAllString(string(*t), rep))
}

// Split splits given string into smaller strings of type Text
func (t *Text) Split(sep string) []Text {
	z := make([]Text, 0)
	x := strings.Split(string(*t), sep)
	for _, v := range x {
		z = append(z, Text(v))
	}
	return z
}

// Substring returns a substring given the start and end indices
func (t *Text) Substring(start, end int) Text {
	s := 0
	length := len(*t)
	e := length - 1
	if start < length {
		s = start
	}
	if end < start {
		e = start
	} else if end < length {
		e = end
	}
	return (*t)[s:e]
}

// Substr return a substring given the start and end indices
func (t *Text) Substr(start, length int) Text {
	s := 0
	slen := len(*t)
	l := slen - 1
	if start < slen {
		s = start
	}

	if (length + s) < slen {
		l = length + s
	}
	return (*t)[s:l]
}

// Length returns the length of the string
func (t *Text) Length() int {
	return len(*t)
}
