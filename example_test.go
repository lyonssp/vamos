package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lyonssp/vamos"
)

func TestTrueProp(t *testing.T) {
	props := vamos.NewProperties(t)
	props.Add("addition is commutative", func(v *vamos.T) {
		a, b := v.Int(), v.Int()
		v.AssertEqual(a+b, b+a)
	})
	tt := &recordingT{}
	props.Test(tt)
	assert.False(t, t.Failed())
}

func TestFalseProp(t *testing.T) {
	props := vamos.NewProperties(t)
	props.Add("true is not false", func(v *vamos.T) {
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
