package vamos

import (
	"math/rand"
)

// batcher manages a batch of property assessments
type batcher struct {
	predicate func(*T)

	// configuration
	n    int   // number of times to execute
	seed int64 // seed used for generators

	// state tracked during execution
	passed int // number of tests runs passed
}

// execute runs property assessments and returns an
// indicator whether they failed or succeeded
func (b *batcher) execute(t TT) bool {
	test := &T{
		t:    t,
		Rand: rand.New(rand.NewSource(b.seed)),
	}

	for i := 0; i < b.n; i++ {
		b.predicate(test)

		if test.failed {
			return false
		}
		b.passed++
	}

	return true
}
