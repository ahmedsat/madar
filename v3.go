package madar

type Vector3 struct {
	X, Y, Z float32
}

// Add implements Vector.
func (v Vector3) Add(v2 Vector) Vector {
	panic("unimplemented")
}

// Distance implements Vector.
func (v Vector3) Distance(v2 Vector) float32 {
	panic("unimplemented")
}

// Divide implements Vector.
func (v Vector3) Divide(v2 Vector) Vector {
	panic("unimplemented")
}

// DivideScalar implements Vector.
func (v Vector3) DivideScalar(s float32) Vector {
	panic("unimplemented")
}

// Dot implements Vector.
func (v Vector3) Dot(v2 Vector) float32 {
	panic("unimplemented")
}

// IsEqual implements Vector.
func (v Vector3) IsEqual(v2 Vector) bool {
	panic("unimplemented")
}

// IsZero implements Vector.
func (v Vector3) IsZero() bool {
	panic("unimplemented")
}

// Length implements Vector.
func (v Vector3) Length() float32 {
	panic("unimplemented")
}

// Lerp implements Vector.
func (v Vector3) Lerp(v2 Vector, t float32) Vector {
	panic("unimplemented")
}

// Multiply implements Vector.
func (v Vector3) Multiply(v2 Vector) Vector {
	panic("unimplemented")
}

// MultiplyScalar implements Vector.
func (v Vector3) MultiplyScalar(s float32) Vector {
	panic("unimplemented")
}

// Negate implements Vector.
func (v Vector3) Negate() Vector {
	panic("unimplemented")
}

// Normalize implements Vector.
func (v Vector3) Normalize() Vector {
	panic("unimplemented")
}

// Rotate implements Vector.
func (v Vector3) Rotate(yaw float32, pitch float32, roll float32) Vector {
	panic("unimplemented")
}

// Scale implements Vector.
func (v Vector3) Scale(s float32) Vector {
	panic("unimplemented")
}

// String implements Vector.
func (v Vector3) String() string {
	panic("unimplemented")
}

// Sub implements Vector.
func (v Vector3) Sub(v2 Vector) Vector {
	panic("unimplemented")
}

var _ Vector = Vector3{}
