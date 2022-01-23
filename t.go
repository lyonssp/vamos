package vamos

import (
	"math/rand"
)

type T struct {
	t TT
	*rand.Rand
	failed bool
}

// supplementary generators

func (test *T) String() string {
	var bs []byte
	_, err := test.Read(bs)
	if err != nil {
		test.failed = true
		return ""
	}
	test.t.Logf("generated string: %s", string(bs))
	return string(bs)
}

func (test *T) Int() int {
	i := test.Int63()
	test.t.Logf("generated int: %d", i)
	return int(i)
}

func (test *T) Intn(n int) int {
	i := test.Int63n(int64(n))
	test.t.Logf("generated int: %d", i)
	return int(i)
}

// testing.T compatibility methods

func (test *T) Errorf(format string, args ...interface{}) {
	test.failed = true
}
