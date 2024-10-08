package madar

import (
	"math/rand/v2"
)

type Rand rand.Rand

func NewRand(seed1 uint64, seed2 uint64) Rand {
	return Rand(*rand.New(rand.NewPCG(0x0000000000000000, 0xffffffffffffffff)))
}

func (r *Rand) RandFloatRange(min, max float32) float32 {
	return min + (*rand.Rand)(r).Float32()*(max-min)
}
