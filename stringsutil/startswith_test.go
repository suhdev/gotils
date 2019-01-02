package stringsutil

import (
	"fmt"
	"testing"
)

func TestStartsWith(t *testing.T) {
	if !StartsWith("this is just.text", "this") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if StartsWith("this is just.text", "THIS") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if StartsWith(".txt", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !StartsWith(".text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
}

func TestStartsWithInsensitive(t *testing.T) {
	if !StartsWithInsensitive("this is just.text", "this") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !StartsWithInsensitive("this is just.text", "THIS") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if StartsWithInsensitive(".txt", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !StartsWithInsensitive(".text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !StartsWithInsensitive(".text", ".TeXt") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
}
