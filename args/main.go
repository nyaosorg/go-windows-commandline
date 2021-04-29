package args

import (
	"strings"
	"unicode"

	"github.com/zetamatta/go-windows-commandline"
)

func SplitArgs(line string) (raw []string, cooked []string) {
	var raw1, cooked1 strings.Builder

	quote := false
	backslash := false

	line = strings.TrimPrefix(line, " ")

	for _, r := range line {
		if backslash {
			backslash = false
			if r != '"' && r != '\\' && r != ' ' {
				cooked1.WriteByte('\\')
			}
		} else {
			if r == '"' {
				quote = !quote
				raw1.WriteByte('"')
				continue
			} else if r == '\\' {
				backslash = true
				raw1.WriteByte('\\')
				continue
			} else if !quote && unicode.IsSpace(r) {
				raw = append(raw, raw1.String())
				cooked = append(cooked, cooked1.String())
				raw1.Reset()
				cooked1.Reset()
				continue
			}
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
