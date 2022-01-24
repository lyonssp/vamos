package vamos

import (
	"math/rand"
)

type Generator func(*rand.Rand) interface{}

func Choice(options ...interface{}) Generator {
	return func(r *rand.Rand) interface{} {
		return options[r.Intn(len(options))]
	}
}
