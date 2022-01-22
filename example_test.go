package vamos

import (
	"fmt"
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

func TestGenString(t *testing.T) {
	props := NewProperties(t)
	props.Add("gen string", func(v *T) {
		s := v.String()
		v.AssertEqual(s, s)
	})
	tt := &recordingT{}
	props.Test(tt)
	assert.False(t, tt.failed)
}

func TestMath(t *testing.T) {
	props := NewProperties(t)
	props.Add("addition is commutative", func(v *T) {
		a := v.Int()
		b := v.Int()
		assert.Equal(v, a+b, b+a)
	})
	props.Test(t)
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
