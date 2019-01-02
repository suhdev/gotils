# gotils 

A set of utility functions and interfaces for go applications 

## dirio 

A set of utility functions to read files in a directory. 

```go

import (
    "fmt"
    "gotils/dirio"
)

cDir := ReadDirChan("/home/suhail",20,nil)

for d := range cDir {
    fmt.Println(d)
}


```