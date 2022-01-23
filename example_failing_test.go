// +build fail

package vamos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailure(t *testing.T) {
	props := NewProperties(t)
	props.Add("a mod b never equals 0", func(v *T) {
		a, b := v.Intn(10), v.Intn(10)
		if a%b != 0 {
			assert.Equal(v, true, false)
		}
	})

	props.Test(t)
}
