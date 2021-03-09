go-windows-commandline
======================

This package gets the raw command-line as is.

```go
[C:] cd example
[C:] type main.go
package main

import (
    "fmt"

    "github.com/zetamatta/go-windows-commandline"
)

func main() {
    cmdline := commandline.Get()
    fmt.Printf("%v\n", cmdline)
}

[C:] go build
[C:] example.exe "1  2" "3  4" "5  6"
example.exe "1  2" "3  4" "5  6"
$
```

Not on Windows-OS, this package emulates the behavior on Windows.
