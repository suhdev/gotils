package stringsutil

import (
	"fmt"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	t1 := "suhail"
	t2 := "sail"
	r := LevenshteinDistance(t1, t2)
	if r != 2 {
		t.Error(fmt.Sprintf("Expected 2 but got %d", r))

	}
	t1 = "Something new"
	t2 = "Something else"
	r = LevenshteinDistance(t1, t2)
	if r != 4 {
		t.Error(fmt.Sprintf("Expected 4 but got %d", r))
	}
}

func TestLevenshteinDistanceWithPath(t *testing.T) {
	t1 := "suhail"
	t2 := "sail"
	r, v := LevenshteinDistanceWithPath(t1, t2)
	if r != 2 || len(v) != 2 {
		t.Error(fmt.Sprintf("Expected 2 but got %d", r))
	}
	t1 = "Something new"
	t2 = "Something exnbew"
	r, m := LevenshteinDistanceWithPath(t1, t2)
	if r == 0 || len(m) != 5 {
		t.Log(m)
		t.Error(fmt.Sprintf("Expected 4 but got %d", r))
	}
}
