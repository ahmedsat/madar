package madar

type Vector interface {
	IsEqual(v2 Vector) bool
	IsZero() bool
	Length() float32
	Scale(s float32) Vector
	Normalize() Vector
	Add(v2 Vector) Vector
	Sub(v2 Vector) Vector
	Multiply(v2 Vector) Vector
	MultiplyScalar(s float32) Vector
	Divide(v2 Vector) Vector
	DivideScalar(s float32) Vector
	Negate() Vector
	Lerp(v2 Vector, t float32) Vector
	Distance(v2 Vector) float32
	Rotate(yaw, pitch, roll float32) Vector
	Dot(v2 Vector) float32
	String() string
}
