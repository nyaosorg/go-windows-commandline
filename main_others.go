// +build !windows

package commandline

import (
	"os"
	"regexp"
	"strings"
)

var rxQuote = regexp.MustCompile(`\\*"`)

func Get() string {
	var buffer strings.Builder
	for i, s := range os.Args {
		if i > 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteByte('"')
		s = rxQuote.ReplaceAllStringFunc(s, func(ss string) string {
			if len(ss) == 1 {
				return `\"`
			}
			return strings.Repeat(`\`, 2*(len(ss)-1)) + `\"`
		})
		buffer.WriteString(s)
		buffer.WriteByte('"')
	}
	return buffer.String()
}
