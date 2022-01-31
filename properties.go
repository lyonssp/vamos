package vamos

import (
	"flag"
	"fmt"
	"math/rand"
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

type Property struct {
	Generate func(*rand.Rand) interface{}
	Simplify func(interface{}) func() (interface{}, bool)
	Check    func(interface{}) bool
}

type GenericProperty[T any] struct {
	Generator GenericGenerator[T]
	Check    func(T) bool
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
// results to *testing.T
func (p *Properties) Test(t *testing.T) {
	t.Helper()

	reporter := testingReporter{t}

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

type testingReporter struct {
	t *testing.T
}

func (dr testingReporter) Report(report PropertyReport) {
	dr.t.Helper()
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

func green(s string) string {
	return fmt.Sprintln(colorGreen, s, colorReset)
}

func red(s string) string {
	return fmt.Sprintln(colorRed, s, colorReset)
}
