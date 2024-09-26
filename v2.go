package madar

import (
	"math"
)

type Vector2 [2]float32

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		v[0] + other[0],
		v[1] + other[1],
	}
}

func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{
		v[0] - other[0],
		v[1] - other[1],
	}
}

func (v Vector2) Multiply(scalar float32) Vector2 {
	return Vector2{
		v[0] * scalar,
		v[1] * scalar,
	}
}

func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1])))
}
