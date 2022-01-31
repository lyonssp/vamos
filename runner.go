package vamos

import (
	"container/list"
	"math/rand"
	"testing"
	"fmt"
	"strings"
	"time"
)

func Check[O any](t *testing.T, p GenericProperty[O]) {
	t.Helper()
	reporter[O]{t}.report(check(p))
}

func check[O any](p GenericProperty[O]) report[O] {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	testcase := p.Generator.Generate(rng)

	if p.Check(testcase) {
		return report[O]{
			failed: false,
			input: testcase,
			minimized: testcase,
		}
	}

	// BFS against tree resulting from shrinks, but
	// short-circuit along any path when the simplified
	// test case does not fail
	ls := list.New()
	ls.PushBack(testcase)

	var simplest O
	for ls.Len() > 0 {
		input := ls.Remove(ls.Front()).(O)
		simplest = input

		shrinker := p.Generator.Simplify(input)
		for {
			next, ok := shrinker()
			if !ok {
				break
			}
			if !p.Check(next) {
				ls.PushBack(next)
				break
			}
		}
	}

	return report[O] {
		failed: true,
		input: testcase,
		minimized: simplest,
	}
}

type report[O any] struct {
	propDesc string
	failed bool
	input O
	minimized O
	numPassed int
	maxChecks int
}

type reporter[O any] struct { t *testing.T }

func (r reporter[O]) report(in report[O]) {
	r.t.Helper()
	if in.failed {
		r.t.Errorf(red(strings.Join([]string{
			"",
			fmt.Sprintf("property: %s", in.propDesc),
			fmt.Sprintf("test case:   %v", in.input),
			fmt.Sprintf("simplified:   %v", in.minimized),
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
