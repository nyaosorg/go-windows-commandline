package args

import (
	"strings"
	"testing"
)

func test(source string, expectRaw, expectCooked []string) (bool, []string, []string) {
	resultRaw, resultCooked := SplitArgs(source)
	if len(expectRaw) != len(resultRaw) {
		return false, resultRaw, resultCooked
	}
	if len(expectCooked) != len(resultCooked) {
		return false, resultRaw, resultCooked
	}
	if len(expectCooked) != len(expectRaw) {
		return false, resultRaw, resultCooked
	}
	for i := range expectRaw {
		if resultRaw[i] != expectRaw[i] {
			return false, resultRaw, resultCooked
		}
		if resultCooked[i] != expectCooked[i] {
			return false, resultRaw, resultCooked
		}
	}
	return true, resultRaw, resultCooked
}

type _TestCase struct {
	source       string
	expectRaw    []string
	expectCooked []string
}

func TestSplitArgs(t *testing.T) {
	var testcases = []_TestCase{
		{
			source:       `ahaha ihihi`,
			expectRaw:    []string{`ahaha`, `ihihi`},
			expectCooked: []string{`ahaha`, `ihihi`},
		},
		{
			source:       `ahaha "C:\Program Files"`,
			expectRaw:    []string{`ahaha`, `"C:\Program Files"`},
			expectCooked: []string{`ahaha`, `C:\Program Files`},
		},
		{
			source:       `ahaha C:\Program\ Files`,
			expectRaw:    []string{`ahaha`, `C:\Program\`, `Files`},
			expectCooked: []string{`ahaha`, `C:\Program\`, `Files`},
		},
		{
			source:       `ahaha \"C:\\Program\ Files\\" a`,
			expectRaw:    []string{`ahaha`, `\"C:\\Program\`, `Files\\" a`},
			expectCooked: []string{`ahaha`, `"C:\\Program\`, `Files\\ a`},
		},
		{
			source:       `-c "ls -a "C:\Program Files""`,
			expectRaw:    []string{`-c`, `"ls -a "C:\Program`, `Files""`},
			expectCooked: []string{`-c`, `ls -a C:\Program`, `Files`},
		},
		{
			source:       ` -c ahaha`,
			expectRaw:    []string{`-c`, `ahaha`},
			expectCooked: []string{`-c`, `ahaha`},
		},
		{
			source:       ` --netuse S:=\\foo\bar\gar`,
			expectRaw:    []string{`--netuse`, `S:=\\foo\bar\gar`},
			expectCooked: []string{`--netuse`, `S:=\\foo\bar\gar`},
		},
		{
			source:       ` --netuse "S:=\\foo\bar\gar"`,
			expectRaw:    []string{`--netuse`, `"S:=\\foo\bar\gar"`},
			expectCooked: []string{`--netuse`, `S:=\\foo\bar\gar`},
		},
	}
	for _, case1 := range testcases {
		if ok, resultRaw, resultCooked := test(case1.source, case1.expectRaw, case1.expectCooked); !ok {
			t.Fatalf("  splitArgs(`%s`)\n!= `%s`/`%s`\n(result=`%s`/`%s`)",
				case1.source,
				strings.Join(case1.expectRaw, "`,`"),
				strings.Join(case1.expectCooked, "`,`"),
				strings.Join(resultRaw, "`,`"),
				strings.Join(resultCooked, "`,`"))
		}
	}
}
