package args

import "testing"

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
			expectRaw:    []string{`ahaha`, `C:\Program\ Files`},
			expectCooked: []string{`ahaha`, `C:\Program Files`},
		},
		{
			source:       `ahaha \"C:\\Program\ Files\\" a`,
			expectRaw:    []string{`ahaha`, `\"C:\\Program\ Files\\" a`},
			expectCooked: []string{`ahaha`, `"C:\Program Files\ a`},
		},
	}
	for _, case1 := range testcases {
		if ok, resultRaw, resultCooked := test(case1.source, case1.expectRaw, case1.expectCooked); !ok {
			t.Fatalf("  splitArgs(`%v`)\n!= %v,%v\n(result=%v,%v)",
				case1.source,
				case1.expectRaw, case1.expectCooked,
				resultRaw, resultCooked)
		}
	}
}
