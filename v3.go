package madar

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float32
}

func (v Vector3) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

func (v Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vector3) Normalize() Vector3 {
	return v.Scale(1 / v.Length())
}

func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vector3) Scale(s float32) Vector3 {
	return Vector3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vector3) Multiply(v2 Vector3) Vector3 {
	return Vector3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vector3) MultiplyScalar(s float32) Vector3 {
	return Vector3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vector3) Divide(v2 Vector3) Vector3 {
	return Vector3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}

func (v Vector3) DivideScalar(s float32) Vector3 {
	return Vector3{v.X / s, v.Y / s, v.Z / s}
}

func (v Vector3) Negate() Vector3 {
	return Vector3{-v.X, -v.Y, -v.Z}
}

func (v Vector3) Lerp(v2 Vector3, t float32) Vector3 {
	return Vector3{
		v.X + (v2.X-v.X)*t,
		v.Y + (v2.Y-v.Y)*t,
		v.Z + (v2.Z-v.Z)*t,
	}
}

func (v Vector3) Distance(v2 Vector3) float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z - v2.X*v2.X - v2.Y*v2.Y - v2.Z*v2.Z)))
}

func (v Vector3) Rotate(yaw, pitch, roll float32) Vector3 {
	rotationMatrix := RotationMatrix4X4(Vector3{yaw, pitch, roll})
	return rotationMatrix.MultiplyVector3(v)
}

func (v Vector3) Dot(v2 Vector3) float32 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v.X, v.Y, v.Z)
}

func (v Vector3) IsEqual(v2 Vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}
