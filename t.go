package vamos

import (
	"fmt"
	"math/rand"
)

// T serves as a test controller and captures metadata for
// a single property assessment
type T struct {
	*rand.Rand
	failed bool
	inputs []interface{}
}

// supplementary generators

func (test *T) String() string {
	var bs []byte
	_, err := test.Read(bs)
	if err != nil {
		test.failed = true
		return ""
	}
	s := fmt.Sprintf("%x", bs)
	test.inputs = append(test.inputs, s)
	return s
}

func (test *T) Intn(n int) int {
	i := test.Rand.Intn(n)
	test.inputs = append(test.inputs, i)
	return i
}

func (test *T) IntRange(a, b int) int {
	if a >= b {
		panic(fmt.Sprintf("cannot generate integer in invalid range [%d,%d)", a, b))
	}
	i := test.Rand.Intn(b-a) + a
	test.inputs = append(test.inputs, i)
	return i
}

func (test *T) Any(fn func(*rand.Rand) interface{}) interface{} {
	out := fn(test.Rand)
	test.inputs = append(test.inputs, out)
	return out
}

// testing.T compatibility methods

func (test *T) Errorf(format string, args ...interface{}) {
	test.failed = true
}
