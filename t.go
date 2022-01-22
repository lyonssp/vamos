package vamos

import (
	"math/rand"
	"reflect"
)

type T struct {
	t TT
	*rand.Rand
	failed bool
}

func (test *T) AssertEqual(x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		test.failed = true
		return
	}
	return
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
	n := test.Int63()
	test.t.Logf("generated int: %d", n)
	return int(n)
}

// testing.T compatibility methods

func (test *T) Errorf(format string, args ...interface{}) {
	test.t.Errorf(format, args...)
}
