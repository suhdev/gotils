package stringsutil

import (
	"fmt"
	"testing"
)

func TestConcat(t *testing.T) {
	x := []string{"1", "2"}
	y := []string{"3", "4"}
	z := Concat(x, y)
	if len(z) != 4 {
		t.Error(fmt.Sprintf("Expected 4 items after concatenation but got %d", len(z)))
	}
	if z[0] != "1" || z[1] != "2" || z[2] != "3" || z[3] != "4" {
		t.Error(fmt.Sprintf("Concatenated items aren't in order expected 1, 2, 3, 4 but got %s", z))
	}

	a := Concat()
	if len(a) > 0 {
		t.Error(fmt.Sprintf("Expected zero array but got %d", len(a)))
	}

}
