package madar

import "math"

type Vector3 [3]float32

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		v[0] + other[0],
		v[1] + other[1],
		v[2] + other[2],
	}
}

func (v Vector3) Subtract(other Vector3) Vector3 {
	return Vector3{
		v[0] - other[0],
		v[1] - other[1],
		v[2] - other[2],
	}
}

func (v Vector3) Multiply(scalar float32) Vector3 {
	return Vector3{
		v[0] * scalar,
		v[1] * scalar,
		v[2] * scalar,
	}
}

func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
}

func (v Vector3) Normalize() Vector3 {
	length := v.Length()
	return Vector3{
		v[0] / length,
		v[1] / length,
		v[2] / length,
	}
}

func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		v[1]*other[2] - v[2]*other[1],
		v[2]*other[0] - v[0]*other[2],
		v[0]*other[1] - v[1]*other[0],
	}
}

func (v Vector3) Dot(other Vector3) float32 {
	return v[0]*other[0] + v[1]*other[1] + v[2]*other[2]
}
