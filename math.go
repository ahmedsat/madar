package madar

import (
	"math"

	"golang.org/x/exp/constraints"
)

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

func RadToDeg(radians float32) float32 {
	return radians * 180 / math.Pi
}

func DegToRad(degrees float32) float32 {
	return degrees * math.Pi / 180
}
