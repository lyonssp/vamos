package vamos

import (
	"fmt"
	"math/rand"
)

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*[]"
)

type Generator func(*rand.Rand) interface{}

type GenericGenerator[T any] interface {
	Generate(r *rand.Rand) T
	Simplify(T) func() (T, bool)
}

type intn struct {
	n int
}

func (i intn) Generate(r *rand.Rand) int {
	return r.Intn(i.n)
}

func (i intn) Simplify(original int) func() (int, bool) {
	half := original / 2

	return func() (int, bool) {
		if half == 0 {
			return 0, false
		}
		out := original - half
		half /= 2
		return out, true
	}
}

func Int() Generator {
	return func(r *rand.Rand) interface{} {
		return r.Int()
	}
}

func Intn(n int) Generator {
	return func(r *rand.Rand) interface{} {
		return r.Intn(n)
	}
}

func IntRange(a, b int) Generator {
	return func(r *rand.Rand) interface{} {
		if a >= b {
			panic(fmt.Sprintf("cannot generate integer in invalid range [%d,%d)", a, b))
		}
		return r.Intn(b-a) + a
	}
}

func AlphabeticString() Generator {
	return genString(uppercase+lowercase, 25)
}

func AlphanumericString() Generator {
	return genString(uppercase+lowercase+digits, 25)
}

func String() Generator {
	return genString(uppercase+lowercase+digits+special, 25)
}

func genString(charset string, n int) Generator {
	return func(r *rand.Rand) interface{} {
		length := r.Intn(n) // [1,n) characters
		out := make([]byte, length)
		for i := 0; i < len(out); i++ {
			out[i] = charset[r.Intn(len(charset))]
		}
		return string(out)
	}
}

func Choice(options ...interface{}) Generator {
	return func(r *rand.Rand) interface{} {
		return options[r.Intn(len(options))]
	}
}
