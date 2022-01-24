package vamos

import (
	"flag"
	"fmt"
	"strings"
	"testing"
	"time"
)

var (
	seedFlag = flag.Int64("vamos.seed", 0, "seed to use to reproduce observed behaviors")

	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type Properties struct {
	seed  int64
	props []property
}

func NewProperties(t *testing.T) *Properties {
	seed := *seedFlag
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	return &Properties{
		seed: seed,
	}
}

type property struct {
	desc      string
	predicate func(t *T)
}

func (p *Properties) Add(desc string, fn func(t *T)) {
	p.props = append(p.props, property{
		desc,
		fn,
	})
}

// Test checks all properties and reports
// results to the given reporter
//
// The builtin *testing.T object can be passed
// as a reporter for typical testing usage
func (p *Properties) Test(t TT) {
	p.Run(defaultReporter{t})
}

// Run checks all properties and reports
// results to the given reporter
//
// The builtin *testing.T object can be passed
// as a reporter for typical testing usage
func (p *Properties) Run(reporter Reporter) {
	n := 100 // number of times to run a check
	for _, prop := range p.props {
		batch := &batcher{
			prop: prop,
			n:    n,
			seed: p.seed,
		}
		reporter.Report(batch.execute())
	}
}

type Reporter interface {
	Report(PropertyReport)
}

type defaultReporter struct {
	t TT
}

func (dr defaultReporter) Report(report PropertyReport) {
	if report.failed {
		dr.t.Errorf(red(strings.Join([]string{
			"",
			fmt.Sprintf("property: %s", report.propDesc),
			fmt.Sprintf("inputs:   %v", report.failureInput),
			fmt.Sprintf("passed:   %d", report.numPassed),
			fmt.Sprintf("max runs: %d", report.maxChecks),
		}, "\n")))
		return
	}

	dr.t.Logf(green(strings.Join([]string{
		"",
		fmt.Sprintf("property: %s", report.propDesc),
		fmt.Sprintf("passed:   %d", report.numPassed),
		fmt.Sprintf("max runs: %d", report.maxChecks),
	}, "\n")))
}

// TT is an interface wrapper around testing.T
type TT interface {
	Errorf(fmt string, args ...interface{})
	Logf(fmt string, args ...interface{})
}

func green(s string) string {
	return fmt.Sprintln(colorGreen, s, colorReset)
}

func red(s string) string {
	return fmt.Sprintln(colorRed, s, colorReset)
}
