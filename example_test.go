package vamos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrueProp(t *testing.T) {
	props := NewProperties(t)
	props.Add("addition is commutative", func(v *T) {
		a, b := v.Int(), v.Int()
		v.AssertEqual(a+b, b+a)
	})
	tt := &recordingT{}
	props.Test(tt)
	assert.False(t, tt.failed)
}

func TestFalseProp(t *testing.T) {
	props := NewProperties(t)
	props.Add("true is not false", func(v *T) {
		v.AssertEqual(true, false)
	})
	tt := &recordingT{}
	props.Test(tt)
	assert.True(t, tt.failed)
}

type recordingT struct {
	failed bool
}

func (r *recordingT) Errorf(_ string, _ ...interface{}) {
	r.failed = true
}
