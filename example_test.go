package vamos

import (
	"fmt"
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

func TestGenString(t *testing.T) {
	props := NewProperties(t)
	props.Add("a string is always equal to itself", func(v *T) {
		s := v.String()
		assert.Equal(v, s, s)
	})
	props.Test(t)
}

func TestFalseProp(t *testing.T) {

	props := NewProperties(t)
	props.Add("true is not false", func(v *T) {
		assert.Equal(v, true, false)
	})
	tt := &recordingT{}
	props.Test(tt)

	assert.True(t, tt.failed)
}

type recordingT struct {
	failed bool
	log    []string
}

func (r *recordingT) Errorf(_ string, _ ...interface{}) {
	r.failed = true
}

func (r *recordingT) Logf(s string, args ...interface{}) {
	r.log = append(r.log, fmt.Sprintf(s, args...))
}

var _ TT = &recordingT{}
