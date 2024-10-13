package madar

import (
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

type Rand rand.Rand

func NewRand(seed1 uint64, seed2 uint64) Rand {
	return Rand(*rand.New(rand.NewPCG(0x0000000000000000, 0xffffffffffffffff)))
}

func (r *Rand) RandFloatRange(min, max float32) float32 {
	return min + (*rand.Rand)(r).Float32()*(max-min)
}

// Clamp is a generic function that constrains a value to lie between two defined bounds.
// It works for any type that supports comparison (integers, floats, etc.).
func Clamp[T constraints.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
