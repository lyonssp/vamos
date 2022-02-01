package vamos

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Check[O any](t *testing.T, p GenericProperty[O]) {
	t.Helper()
	reporter[O]{t}.report(check(p))
}

func check[O any](p GenericProperty[O]) report[O] {
	n := 100 // number of runs
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		_testcase := p.Generator.Generate(rng)

		if !p.Check(_testcase) {
			return report[O]{
				propDesc:   "placeholder",
				failed:     true,
				testcase:   _testcase,
				simplified: simplify(_testcase, p),
				numPassed:  i,
				maxChecks:  100,
			}
		}

	}

	return report[O]{
		propDesc:  "placeholder",
		failed:    false,
		numPassed: n,
		maxChecks: 100,
	}
}

// BFS against tree resulting from shrinks, but
// short-circuit along any path when the simplified
// test case does not fail
func simplify[O any](testcase O, prop GenericProperty[O]) O {
	ls := list.New()
	ls.PushBack(testcase)

	var simplest O
	for ls.Len() > 0 {
		simplest = ls.Remove(ls.Front()).(O)
		shrinker := prop.Generator.Simplify(simplest)
		for {
			next, ok := shrinker()
			if !ok {
				break
			}
			if !prop.Check(next) {
				ls.PushBack(next)
				break
			}
		}
	}

	return simplest
}

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
