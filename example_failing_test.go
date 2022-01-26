// +build fail

package vamos

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailure(t *testing.T) {
	props := NewProperties(t)

	props.Add("true is false", func(v *T) {
		assert.True(v, false)
	})

	props.Add("a mod b never equals 0", func(v *T) {
		a, b := v.IntRange(1, 10), v.IntRange(1, 10)
		assert.NotEqual(v, a%b, 0)
	})

	props.Add("no two people are the same", func(v *T) {
		x := v.Any(genPerson).(person)
		y := v.Any(genPerson).(person)
		assert.NotEqual(v, x, y)
	})

	props.Test(t)
}

type person struct {
	firstName  string
	middleName string
	homeState  string
}

func genPerson(rng *rand.Rand) interface{} {
	return person{
		firstName:  Choice("Bob", "Chris", "Sean")(rng).(string),
		middleName: Choice("Patrick", "Mary", "Louie")(rng).(string),
		homeState:  Choice("Virginia", "Texas", "New York")(rng).(string),
	}
}
