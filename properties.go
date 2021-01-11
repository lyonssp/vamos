package vamos

import (
	"math/rand"
	"testing"
	"time"
)

type Properties struct {
	seed  int64
	props []property
}

func NewProperties(t *testing.T) *Properties {
	return &Properties{
		seed: time.Now().UnixNano(),
	}
}

type property struct {
	desc      string
	predicate func(t *T)
}

func (p *Properties) Add(desc string, fn func(t *T)) {
	p.props = append(p.props, property{
		desc,
		fn,
	})
}

func (p *Properties) Test(t TT) {
	var failures []property
	for _, prop := range p.props {
		test := &T{Rand: rand.New(rand.NewSource(p.seed))}
		prop.predicate(test) // TODO: wrap inside some executor for customization (e.g. number of runs)
		if test.failed {
			failures = append(failures, prop)
			continue
		}
	}

	if len(failures) > 0 {
		t.Errorf("%d/%d properties did not hold", len(failures), len(p.props))
	}
}

// TT is an interface wrapper around testing.T
type TT interface {
	Errorf(msg string, args ...interface{})
}
