package vamos

import (
	"flag"
	"math/rand"
	"testing"
	"time"
)

var seedFlag = flag.Int64("vamos.seed", 0, "seed to use to reproduce observed behaviors")

type Properties struct {
	seed  int64
	props []property
}

func NewProperties(t *testing.T) *Properties {
	seed := *seedFlag
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	return &Properties{
		seed: seed,
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
	t.Logf("seed: %d", p.seed)

	var failures []property
	for _, prop := range p.props {
		test := &T{
			t:    t,
			Rand: rand.New(rand.NewSource(p.seed)),
		}
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
	Logf(fmt string, args ...interface{})
}
