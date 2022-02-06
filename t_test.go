package vamos

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueProp(t *testing.T) {
	Check(t, Property[Pair[int]]{
		"addition is commutative",

		PairGenerator(IntRange(0, 100)),

		func(v *V[Pair[int]]) {
			input := v.Input
			assert.Equal(v, input.Left+input.Right, input.Right+input.Left)
		},
	})
}

func TestInt(t *testing.T) {
	Check(t, Property[int]{
		"Int always generates integers within max bounds",

		Int(),

		func(v *V[int]) {
			x := v.Input
			assert.GreaterOrEqual(v, x, math.MinInt)
			assert.LessOrEqual(v, x, math.MaxInt)
		},
	})
}

func TestIntn(t *testing.T) {
	Check(t, Property[int]{
		"Intn always generates integers within bounds",

		Intn(100),

		func(v *V[int]) {
			x := v.Input
			assert.GreaterOrEqual(v, x, 0)
			assert.LessOrEqual(v, x, 100)
		},
	})
}

func TestIntRange(t *testing.T) {
	Check(t, Property[int]{
		"IntRange always generates integers within bounds",

		IntRange(-100, 100),

		func(v *V[int]) {
			x := v.Input
			assert.GreaterOrEqual(v, x, -100)
			assert.LessOrEqual(v, x, 100)
		},
	})
}

func TestInvalidIntRange(t *testing.T) {
	assert.Panics(t, func() {
		Check(t, Property[int]{
			"bad generator",

			IntRange(100, -100),

			func(v *V[int]) {
				assert.True(v, true) // should not get here
			},
		})
	})
}

func TestString(t *testing.T) {
	Check(t, Property[string]{
		"strings always equal themselves",

		String(),

		func(v *V[string]) {
			assert.Equal(v, v.Input, v.Input)
		},
	})
}

func TestAlphabeticString(t *testing.T) {
	Check(t, Property[string]{
		"alphabetic strings always equal themselves",

		AlphabeticString(),

		func(v *V[string]) {
			assert.Equal(v, v.Input, v.Input)
		},
	})
}

func TestAlphanumericString(t *testing.T) {
	Check(t, Property[string]{
		"alphabetic strings always equal themselves",

		AlphanumericString(),

		func(v *V[string]) {
			assert.Equal(v, v.Input, v.Input)
		},
	})
}

func TestChoice(t *testing.T) {
	Check(t, Property[string]{
		"choosing from one options always generates that option",

		Choice("foo"),

		func(v *V[string]) {
			assert.Equal(v, v.Input, "foo")
		},
	})

	Check(t, Property[string]{
		"choosing from a list of options always generates one of the options",

		Choice("foo", "bar", "baz"),

		func(v *V[string]) {
			assert.Contains(v, []string{"foo", "bar", "baz"}, v.Input)
		},
	})
}
