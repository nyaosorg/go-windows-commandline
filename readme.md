go-windows-commandline
======================

This package gets the raw command-line as is.

```go
$ cd example
$ type main.go
package main

import (
    "fmt"

    "github.com/zetamatta/go-windows-commandline"
)

func main() {
    cmdline := commandline.Get()
    fmt.Printf("%v\n", cmdline)
}

$ go build
$ example.exe "1  2" "3  4" "5  6"
example.exe "1  2" "3  4" "5  6"
$
```
