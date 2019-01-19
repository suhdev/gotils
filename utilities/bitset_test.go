package utilities_test

import (
	"fmt"
	"gotils/utilities"
	"testing"
)

func TestSet(t *testing.T) {
	var set = utilities.NewBitSet()
	set.Set(10)
	if set.Size() != 2 {
		t.Fatalf("Size is expected to be 2 bytes but got %d", set.Size())
	}
	if !set.Get(10) {
		t.Fatal("Bit no. 10 is expected to be set (1)")
	}

	if fmt.Sprintf("%s", set) != "{10000000000 2}" {
		t.Fatalf("Expected {1000000000 2} but got %s", set)
	}
}
