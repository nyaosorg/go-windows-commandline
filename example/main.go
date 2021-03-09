package main

import (
	"fmt"

	"github.com/zetamatta/go-windows-commandline"
)

func main() {
	cmdline := commandline.Get()
	fmt.Printf("%v\n", cmdline)
}
