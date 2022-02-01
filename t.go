package vamos

import (
	"math/rand"
)

// T serves as a test controller and captures metadata for
// a single property assessment
type T struct {
	*rand.Rand
	failed bool
}

// supplementary generators

/*
func (test *T) String() string {
	return test.gen(String()).(string)
}

func (test *T) AlphabeticString() string {
	return test.gen(AlphabeticString()).(string)
}

func (test *T) AlphanumericString() string {
	return test.gen(AlphanumericString()).(string)
}

func (test *T) Int() int {
	return test.gen(Int()).(int)
}

func (test *T) Intn(n int) int {
	return test.gen(Intn(n)).(int)
}

func (test *T) IntRange(a, b int) int {
	return test.gen(IntRange(a, b)).(int)
}

// gen uses the given generator to generate data and
// then records that data to the test input audit trail
func (test *T) gen(g Generator[O any]) O {
	out := g(test.Rand)
	return out
}
*/

// testing.T compatibility methods

func (test *T) Errorf(format string, args ...interface{}) {
	test.failed = true
}
