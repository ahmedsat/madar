package madar

import "math"

type Vector4 [4]float32

func (v Vector4) Add(other Vector4) Vector4 {
	return Vector4{
		v[0] + other[0],
		v[1] + other[1],
		v[2] + other[2],
		v[3] + other[3],
	}
}

func (v Vector4) Subtract(other Vector4) Vector4 {
	return Vector4{
		v[0] - other[0],
		v[1] - other[1],
		v[2] - other[2],
		v[3] - other[3],
	}
}

func (v Vector4) Multiply(scalar float32) Vector4 {
	return Vector4{
		v[0] * scalar,
		v[1] * scalar,
		v[2] * scalar,
		v[3] * scalar,
	}
}

func (v Vector4) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3])))
}
