package stringsutil

import (
	"fmt"
	"testing"
)

func TestTextSubstr(t *testing.T) {
	x := Text("This is just a test")
	v := x.Substr(1, 3)
	if v != x[1:4] {
		t.Error(fmt.Sprintf("Expected %s but got %s", x[1:4], v))
	}
	v = x.Substr(1, 1)
	if v != Text(x[1]) {
		t.Error(fmt.Sprintf("Expected %s but got %s", Text(x[1]), v))
	}
	v = x.Substr(1, 100)
	if v != x[1:len(x)-1] {
		t.Error(fmt.Sprintf("Expected %s but got %s", x[1:len(x)-1], v))
	}
	v = x.Substr(5, 2)
	if v != x[5:7] {
		t.Error(fmt.Sprintf("Expected %s but got %s", x[5:7], v))
	}
	return
}
