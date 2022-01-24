package vamos

import (
	"math/rand"
)

// batcher manages a batch of property assessments
type batcher struct {
	prop property

	// configuration
	n    int   // number of times to execute
	seed int64 // seed used for generators
}

// execute runs property assessments and returns an
// indicator whether they failed or succeeded
func (b *batcher) execute() PropertyReport {
	rng := rand.New(rand.NewSource(b.seed))

	for i := 0; i < b.n; i++ {
		test := &T{Rand: rng}

		b.prop.predicate(test)

		if test.failed {
			return PropertyReport{
				propDesc:     b.prop.desc,
				numPassed:    i,
				maxChecks:    b.n,
				failed:       true,
				failureInput: test.inputs,
				seed:         b.seed,
			}
		}
	}

	return PropertyReport{
		propDesc:  b.prop.desc,
		numPassed: b.n,
		maxChecks: b.n,
		seed:      b.seed,
	}
}

type PropertyReport struct {
	propDesc     string
	failed       bool
	numPassed    int
	maxChecks    int
	failureInput []interface{}
	seed         int64
}
