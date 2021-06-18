package args

import (
	"strings"
	"unicode"

	"github.com/zetamatta/go-windows-commandline"
)

func SplitArgs(line string) (raw []string, cooked []string) {
	var raw1, cooked1 strings.Builder

	quote := false
	backslashCount := 0

	for _, r := range line {
		if r == '"' {
			if backslashCount%2 == 1 { // \"
				for n := (backslashCount - 1) / 2; n > 0; n-- {
					cooked1.WriteByte('\\')
				}
				backslashCount = 0
				cooked1.WriteByte('"')
			} else {
				for ; backslashCount > 0; backslashCount-- {
					cooked1.WriteByte('\\')
				}
				quote = !quote
			}
			raw1.WriteByte('"')
			continue
		}
		if r == '\\' {
			backslashCount++
			raw1.WriteByte('\\')
			continue
		}
		for ; backslashCount > 0; backslashCount-- {
			cooked1.WriteByte('\\')
		}
		if !quote && unicode.IsSpace(r) {
			if raw1.Len() > 0 {
				raw = append(raw, raw1.String())
				cooked = append(cooked, cooked1.String())
			}
			raw1.Reset()
			cooked1.Reset()
			continue
		}
		raw1.WriteRune(r)
		cooked1.WriteRune(r)
	}
	if raw1.Len() > 0 {
		raw = append(raw, raw1.String())
		cooked = append(cooked, cooked1.String())
	}
	return
}

func Parse() (raw []string, cooked []string) {
	return SplitArgs(commandline.Get())
}
