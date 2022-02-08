//go:build fail
// +build fail

package vamos

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestFailureCheck(t *testing.T) {
	Check(t, Property[bool]{
		"all booleans are true",

		Choice(true, false),

		func(v *V, b bool) {
			assert.True(v, b)
		},
	})
}

func TestStringCheck(t *testing.T) {
	Check(t, Property[string]{
		"strings are at most 2 characters long",

		String(),

		func(v *V, s string) {
			assert.True(v, len(s) <= 2)
		},
	})
}

func TestPairCheck(t *testing.T) {
	Check(t, Property[Pair[int]]{
		"Pair[int].Left < Pair[int].Right",

		PairGenerator(IntRange(0, 100)),

		func(v *V, p Pair[int]) {
			assert.Less(v, p.Left, p.Right)
		},
	})

	Check(t, Property[Pair[int]]{
		"Pair[int].Left >= Pair[int].Right",

		PairGenerator(IntRange(0, 100)),

		func(v *V, p Pair[int]) {
			assert.GreaterOrEqual(v, p.Left, p.Right)
		},
	})

	Check(t, Property[Pair[string]]{
		"strings are all the same length",

		PairGenerator(String()),

		func(v *V, p Pair[string]) {
			assert.Equal(v, len(p.Left), len(p.Right))
		},
	})
}

type person struct {
	firstName  string
	middleName string
	homeState  string
}

func genPerson(rng *rand.Rand) interface{} {
	return person{
		firstName:  Choice("Bob", "Chris", "Sean").Generate(rng),
		middleName: Choice("Patrick", "Mary", "Louie").Generate(rng),
		homeState:  Choice("Virginia", "Texas", "New York").Generate(rng),
	}
}
