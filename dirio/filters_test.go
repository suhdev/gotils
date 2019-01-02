package dirio

import (
	"fmt"
	"testing"
)

func TestFilterJectFiles(t *testing.T) {
	if !FilterJectFiles("file.ject") {
		t.Error(fmt.Sprintf("Expected true but got false"))
	}
}
