package vamos

import (
	"flag"
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

	for _, prop := range p.props {

		batch := &batcher{
			predicate: prop.predicate,
			n:         100,
			seed:      p.seed,
		}
		pass := batch.execute(t) // TODO: wrap inside some executor for customization (e.g. number of runs)

		if !pass {
			t.Errorf(`
			property: %s
			passed:   %d
			max runs: %d
			`, prop.desc, batch.passed, batch.n)
		}
	}
}

// TT is an interface wrapper around testing.T
type TT interface {
	Errorf(fmt string, args ...interface{})
	Logf(fmt string, args ...interface{})
}
