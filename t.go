package vamos

import (
	"math/rand"
	"reflect"
)

type T struct {
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
