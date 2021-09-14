go-windows-commandline
======================

This package gets the raw command-line as is.

```go
[C:] cd example
[C:] type main.go
package main

import (
    "fmt"

    "github.com/nyaosorg/go-windows-commandline"
)

func main() {
    fmt.Println(commandline.Get())
}

[C:] go build
[C:] example.exe "1  2" "3  4" "5  6"
example.exe "1  2" "3  4" "5  6"
$
```

Not on Windows-OS, this package emulates the behavior on Windows.
