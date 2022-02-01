package vamos

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueProp(t *testing.T) {
	Check(t, GenericProperty[Pair[int]]{
		Generator: PairGenerator(IntRange(0, 100)),
		Check: func(integers Pair[int]) bool {
			return integers.Left+integers.Right == integers.Right+integers.Left
		},
	})
}

func TestInt(t *testing.T) {
	Check(t, GenericProperty[int]{
		Generator: Int(),
		Check: func(x int) bool {
			return x >= math.MinInt && x <= math.MaxInt
		},
	})
}

func TestIntn(t *testing.T) {
	Check(t, GenericProperty[int]{
		Generator: Intn(100),
		Check: func(x int) bool {
			return x >= 0 && x <= 100
		},
	})
}

func TestIntRange(t *testing.T) {
	Check(t, GenericProperty[int]{
		Generator: IntRange(-100, 100),
		Check: func(x int) bool {
			return x >= -100 && x <= 100
		},
	})
}

func TestInvalidIntRange(t *testing.T) {
	assert.Panics(t, func() {
		Check(t, GenericProperty[int]{
			Generator: IntRange(100, -100),
			Check: func(x int) bool {
				return true // should not get here
			},
		})
	})
}

func TestString(t *testing.T) {
	Check(t, GenericProperty[string]{
		Generator: String(),
		Check: func(x string) bool {
			return x == x
		},
	})
}

func TestAlphabeticString(t *testing.T) {
	Check(t, GenericProperty[string]{
		Generator: String(),
		Check: func(x string) bool {
			return x == x
		},
	})
}

func TestAlphanumericString(t *testing.T) {
	Check(t, GenericProperty[string]{
		Generator: String(),
		Check: func(x string) bool {
			return x == x
		},
	})
}

func TestChoice(t *testing.T) {
	Check(t, GenericProperty[string]{
		Generator: Choice("foo"),
		Check: func(s string) bool {
			return s == "foo"
		},
	})

	Check(t, GenericProperty[string]{
		Generator: Choice("foo", "bar", "baz"),
		Check: func(s string) bool {
			contains := func(ls []string, s string) bool {
				for _, e := range ls {
					if s == e {
						return true
					}
				}
				return false
			}
			return contains([]string{"foo", "bar", "baz"}, s)
		},
	})
}
