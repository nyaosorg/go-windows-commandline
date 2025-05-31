go-windows-commandline
======================

[![Go Reference](https://pkg.go.dev/badge/github.com/nyaosorg/go-windows-commandline.svg)](https://pkg.go.dev/github.com/nyaosorg/go-windows-commandline)

This Go package provides access to the **raw command-line string** as passed to a process on **Windows**, *before* it is split into `os.Args`.

Why?
----

In Go, `os.Args` gives you the parsed list of command-line arguments. This works well on Unix-like systems because the OS kernel passes arguments as an array of strings. However, on **Windows**, the situation is different:

* The Windows API passes command-line arguments to a process as **a single UTF-16 encoded string**.
* The Go runtime then parses this raw string into a slice of strings, following specific parsing rules.
* This behavior may lead to subtle discrepancies, especially in how quotation marks and whitespace are handled.

For applications that need to:

* Reconstruct the original command-line input,
* Handle custom quoting or parsing logic,
* Perform low-level debugging or shell-related tasks,

retrieving the **unparsed command-line string** is essential.

This package gives you that raw string.

Example
-------

```go
package main

import (
    "fmt"
    "github.com/nyaosorg/go-windows-commandline"
)

func main() {
    fmt.Println(commandline.Get())
}
```

Assuming your source is saved as `main.go`:

```powershell
C:\> go build
C:\> example.exe "1  2" "3  4" "5  6"
example.exe "1  2" "3  4" "5  6"
```

Instead of splitting the arguments, this package returns the full raw command-line string as it was passed to the program.

Cross-Platform Note
-------------------

On non-Windows platforms, this package emulates Windows-like behavior by reconstructing a command-line string from `os.Args`. This allows you to write cross-platform code or test Windows-specific logic on other OSes.

Author
------

[hymkor (HAYAMA Kaoru)](https://github.com/hymkor)

License
-------

MIT License
