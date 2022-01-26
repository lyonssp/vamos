package vamos

import (
	"math"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueProp(t *testing.T) {
	props := NewProperties(t)
	props.Add("addition is commutative", func(v *T) {
		a, b := v.IntRange(0, 100), v.IntRange(0, 100)
		assert.Equal(v, a+b, b+a)
	})
	props.Test(t)
}

func TestInt(t *testing.T) {
	props := NewProperties(t)
	props.Add("generated integer is always within the specified bounds", func(v *T) {
		x := v.Int()
		assert.True(v, x >= math.MinInt)
		assert.True(v, x <= math.MaxInt)
	})
	props.Test(t)
}

func TestIntn(t *testing.T) {
	props := NewProperties(t)
	props.Add("generated integer is always within the specified bounds", func(v *T) {
		x := v.Intn(100)
		assert.True(v, x >= 0)
		assert.True(v, x < 100)
	})
	props.Test(t)
}

func TestIntRange(t *testing.T) {
	props := NewProperties(t)

	props.Add("generated integer is always within the specified range", func(v *T) {
		x := v.IntRange(-100, 100)
		assert.True(v, x >= -100)
		assert.True(v, x < 100)
	})

	props.Test(t)
}

func TestInvalidIntRange(t *testing.T) {

	assert.Panics(t, func() {
		props := NewProperties(t)

		props.Add("invalid int ranges cause panics", func(v *T) {
			a := v.Int()               // [0, )
			b := v.IntRange(-100, a+1) // [-100, a]
			v.IntRange(a, b)           // invalid range
		})

		props.Test(t)
	})
}

func TestString(t *testing.T) {
	props := NewProperties(t)
	props.Add("a string is always equal to itself", func(v *T) {
		s := v.String()
		assert.Equal(v, s, s)
	})
	props.Test(t)
}

func TestAlphabeticString(t *testing.T) {
	props := NewProperties(t)
	props.Add("a generated alphabetic string contains appropriate characters", func(v *T) {
		s := v.AlphabeticString()

		re := regexp.MustCompile("^[a-zA-Z]*$")
		assert.True(v, re.MatchString(s))
	})
	props.Test(t)
}

func TestAlphanumericString(t *testing.T) {
	props := NewProperties(t)
	props.Add("a generated alphanumeric string contains appropriate characters", func(v *T) {
		s := v.AlphanumericString()

		re := regexp.MustCompile("^[a-zA-Z0-9]*$")
		assert.True(v, re.MatchString(s))
	})
	props.Test(t)
}

func TestChoice(t *testing.T) {
	props := NewProperties(t)

	props.Add("generated result is always the same in singleton list of choices", func(v *T) {
		s := v.Any(Choice("foo")).(string)
		assert.Equal(v, "foo", s)
	})

	props.Add("generated result is always in list of choices", func(v *T) {
		choices := []interface{}{"foo", "bar", "baz"}
		s := v.Any(Choice(choices...)).(string)
		assert.Contains(v, choices, s)
	})
	props.Test(t)
}
