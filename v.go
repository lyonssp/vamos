package vamos

import (
	"math/rand"
)

// V serves as a test controller and captures metadata for
// a single property assessment
type V struct {
	rng    *rand.Rand
	failed bool
}

// testing.T compatibility methods

func (v *V) Errorf(format string, args ...interface{}) {
	v.failed = true
}
