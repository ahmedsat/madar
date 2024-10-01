package madar

import "math/rand/v2"

func RandFloatRange(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
