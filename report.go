package vamos

import (
	"fmt"
	"strings"
	"testing"
)

var (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type report[O any] struct {
	propDesc   string
	failed     bool
	testcase   O
	simplified O
	numPassed  int
	maxChecks  int
}

type reporter[O any] struct{ t *testing.T }

func (r reporter[O]) report(in report[O]) {
	r.t.Helper()
	if in.failed {
		r.t.Errorf(red(strings.Join([]string{
			"",
			fmt.Sprintf("property: %s", in.propDesc),
			fmt.Sprintf("test case:   %v", in.testcase),
			fmt.Sprintf("simplified:   %v", in.simplified),
			fmt.Sprintf("passed:   %d", in.numPassed),
			fmt.Sprintf("max runs: %d", in.maxChecks),
		}, "\n")))
		return
	}

	r.t.Logf(green(strings.Join([]string{
		"",
		fmt.Sprintf("property: %s", in.propDesc),
		fmt.Sprintf("passed:   %d", in.numPassed),
		fmt.Sprintf("max runs: %d", in.maxChecks),
	}, "\n")))
}

func green(s string) string {
	return fmt.Sprintln(colorGreen, s, colorReset)
}

func red(s string) string {
	return fmt.Sprintln(colorRed, s, colorReset)
}
