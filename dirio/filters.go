package dirio

import (
	"gotils/stringsutil"
)

// FilterJectFiles filters ject files
// This method is case insensitive
func FilterJectFiles(filename string) bool {
	return stringsutil.EndsWithInsensitive(filename, ".ject")
}
