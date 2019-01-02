package stringsutil

import (
	"fmt"
	"testing"
)

func TestEndsWith(t *testing.T) {
	if !EndsWith("this is just.text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if EndsWith("this is just.text", ".TEXT") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if EndsWith(".txt", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !EndsWith(".text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
}

func TestEndsWithInsensitive(t *testing.T) {
	if !EndsWithInsensitive("this is just.text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !EndsWithInsensitive("this is just.text", ".TEXT") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if EndsWithInsensitive(".txt", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !EndsWithInsensitive(".text", ".text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
	if !EndsWithInsensitive(".text", ".Text") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
}
