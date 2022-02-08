package vamos

import (
	"container/list"
	"flag"
	"math/rand"
	"testing"
	"time"
)

var (
	seedFlag = flag.Int64("vamos.seed", 0, "seed to use to reproduce observed behaviors")
)

type Property[O any] struct {
	Description string
	Generator   Generator[O]
	Check       func(*V, O)
}

func Check[O any](t *testing.T, p Property[O]) {
	t.Helper()

	r := reporter[O]{t}

	r.report(check(t, p))
}

func check[O any](t *testing.T, p Property[O]) report[O] {
	n := 100 // number of runs

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		_testcase := p.Generator.Generate(rng)
		v := &V{rng: rng}
		p.Check(v, _testcase)

		if v.failed {
			return report[O]{
				propDesc:   p.Description,
				failed:     true,
				testcase:   _testcase,
				simplified: simplify(v, _testcase, p),
				numPassed:  i,
				maxChecks:  100,
			}
		}
	}

	return report[O]{
		propDesc:  p.Description,
		failed:    false,
		numPassed: n,
		maxChecks: 100,
	}
}

// BFS against tree resulting from shrinks, but
// short-circuit along any path when the simplified
// test case does not fail
func simplify[O any](v *V, testcase O, prop Property[O]) O {
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

			_v := &V{rng: v.rng}
			prop.Check(_v, next)

			if _v.failed {
				ls.PushBack(next)
				break
			}
		}
	}

	return simplest
}
