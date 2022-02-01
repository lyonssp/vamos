package vamos

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*[]"
)

type Generator[T any] interface {
	Generate(r *rand.Rand) T
	Simplify(T) func() (T, bool)
}

func Int() Generator[int] {
	return IntRange(math.MinInt, math.MaxInt)
}

func Intn(n int) Generator[int] {
	return IntRange(0, n)
}

func IntRange(a, b int) Generator[int] {
	if a >= b {
		panic(fmt.Sprintf("cannot generate integer in invalid range [%d,%d)", a, b))
	}
	return intRange{a, b}
}

type intRange struct {
	min int
	max int
}

func (i intRange) Generate(r *rand.Rand) int {
	if i.min == math.MinInt && i.max == math.MaxInt {
		return r.Int()
	}
	return r.Intn(i.max-i.min) + i.min
}

func (i intRange) Simplify(original int) func() (int, bool) {
	return simplifyInt(i.min, original)
}

func AlphabeticString() Generator[string] {
	return str{lowercase + uppercase, 25}
}

func AlphanumericString() Generator[string] {
	return str{lowercase + uppercase + digits, 25}
}

func String() Generator[string] {
	return str{lowercase + uppercase + digits + special, 25}
}

// str generates a string from the given characterset
// with length in the range [0, str.length]
//
// values simplify towards shorter strings and invdividual
// characters simplify toward the front of the character set string
type str struct {
	charset string
	length  int
}

func (s str) Generate(r *rand.Rand) string {
	length := r.Intn(s.length) // [1,n) characters
	out := make([]byte, length)
	for i := 0; i < len(out); i++ {
		out[i] = s.charset[r.Intn(len(s.charset))]
	}
	return string(out)
}

func (s str) Simplify(original string) func() (string, bool) {
	return simplifyString(original, s.charset)
}

// shrink string size by reducing the length 
// of the input string by factors of two
func shrinkSize(original, charset string) func() (string, bool) {
	next := simplifyInt(0, len(original))
	return func() (string, bool) {
		i, ok := next()
		if !ok {
			return "", false
		}
		if i == 0 {
			return "", true
		}
		return original[:i], true
	}
}

// shrink characters in string starting with the first character in the string
func shrinkCharacters(original, charset string) func() (string, bool) {
	if len(original) <= 1 {
		return noopSimplifier[string]
	}

	at := 0
	return func() (string, bool) {
		for i := at; i < len(original); i++ {
			str, ok := simplifyCharAt(at, original, charset)()
			if !ok {
				at++
				continue
			}

			return str, true
		}

		return "", false
	}
}

func Choice[O any](opts ...O) Generator[O] {
	return choice[O]{options: opts}
}

type choice[O any] struct {
	options []O
}

func (c choice[O]) Generate(r *rand.Rand) O {
	return c.options[r.Intn(len(c.options))]
}

func (c choice[O]) Simplify(original O) func() (O, bool) {
	return func() (O, bool) {
		return original, false
	}
}

func PairGenerator[O any](g Generator[O]) Generator[Pair[O]] {
	return pairGenerator[O]{g}
}

type pairGenerator[O any] struct {
	inner Generator[O]
}

func (p pairGenerator[O]) Generate(r *rand.Rand) Pair[O] {
	return Pair[O]{
		Left:  p.inner.Generate(r),
		Right: p.inner.Generate(r),
	}
}

func (c pairGenerator[O]) Simplify(original Pair[O]) func() (Pair[O], bool) {
	left := original.Left
	right := original.Right
	leftSimplify := c.inner.Simplify(original.Left)
	rightSimplify := c.inner.Simplify(original.Right)
	return func() (Pair[O], bool) {
		// simplify pair.left as much as possible first
		simpleLeft, leftOk := leftSimplify()
		if leftOk {
			left = simpleLeft
		}

		// simplify pair.right next
		simpleRight, rightOk := rightSimplify()
		if rightOk {
			right = simpleRight
		}

		return Pair[O]{left, right}, leftOk || rightOk
	}
}

type Pair[O any] struct {
	Left, Right O
}

type concatter[O any] struct {
	cur    int
	stream []func() (O, bool)
}

func (c concatter[O]) next() (out O, ok bool) {
	for c.cur < len(c.stream) {
		out, ok = c.stream[c.cur]()
		if !ok {
			c.cur++
			continue
		}
		return
	}
	return
}

func concat[O any](simplifiers ...func() (O, bool)) concatter[O] {
	return concatter[O]{
		stream: simplifiers,
	}
}

func noopSimplifier[O any]() (O, bool) {
	var out O
	return out, false
}

func findChar(s string, b byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return i
		}
	}
	return -1
}

func simplifyInt(base, x int) func() (int, bool) {
	if x == base {
		return noopSimplifier[int]
	}

	var exhausted bool
	return func() (int, bool) {
		if exhausted {
			return 0, false
		}
		if x == base {
			exhausted = true
		}
		x = base + (x - base) / 2
		return x, true
	}
}

func simplifyString(original, charset string) func() (string, bool) {
	return concat(
		shrinkSize(original, charset),
		shrinkCharacters(original, charset),
	).next
}

func simplifyCharAt(at int, str, charset string) func() (string, bool) {
	if at < 0 || at >= len(str) {
		return noopSimplifier[string]
	}

	pos := findChar(charset, str[at])
	next := simplifyInt(0, pos)
	return func() (string, bool) {
		i, ok := next()
		if !ok {
			return "", false
		}

		out := make([]byte, len(str))
		copy(out[:at], str[:at])
		out[at] = charset[i]
		copy(out[at+1:], str[at+1:])

		return string(out), true
	}
}
