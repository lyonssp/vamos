//go:build fail
// +build fail

package vamos

import (
	"math/rand"
	"testing"
)

func TestFailureCheck(t *testing.T) {
	Check(t, GenericProperty[bool]{
		Generator: Choice(true, false),
		Check: func(b bool) bool {
			return b
		},
	})
}

func TestStringCheck(t *testing.T) {
	Check(t, GenericProperty[string]{
		Generator: String(),
		Check: func(s string) bool {
			return len(s) <= 2
		},
	})
}

func TestPairCheck(t *testing.T) {
	Check(t, GenericProperty[Pair[int]]{
		Generator: PairGenerator(IntRange(0, 100)),
		Check: func(p Pair[int]) bool {
			return p.Left < p.Right
		},
	})

	Check(t, GenericProperty[Pair[int]]{
		Generator: PairGenerator(IntRange(0, 100)),
		Check: func(p Pair[int]) bool {
			return p.Left >= p.Right
		},
	})

	Check(t, GenericProperty[Pair[string]]{
		Generator: PairGenerator(String()),
		Check: func(p Pair[string]) bool {
			return len(p.Left) < len(p.Right)
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
