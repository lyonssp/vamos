package vamos

import (
	"testing"
)

func TestCheck(t *testing.T) {
	Check(t, GenericProperty[int]{
		Generator: 	intn{100},
		Check: func(x int) bool { return x < 3 },
	})
}
