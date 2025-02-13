package madar

type Vector4 struct {
	X, Y, Z, W float32
}

// Add implements Vector.
func (v Vector4) Add(v2 Vector) Vector {
	panic("unimplemented")
}

// Distance implements Vector.
func (v Vector4) Distance(v2 Vector) float32 {
	panic("unimplemented")
}

// Divide implements Vector.
func (v Vector4) Divide(v2 Vector) Vector {
	panic("unimplemented")
}

// DivideScalar implements Vector.
func (v Vector4) DivideScalar(s float32) Vector {
	panic("unimplemented")
}

// Dot implements Vector.
func (v Vector4) Dot(v2 Vector) float32 {
	panic("unimplemented")
}

// IsEqual implements Vector.
func (v Vector4) IsEqual(v2 Vector) bool {
	panic("unimplemented")
}

// IsZero implements Vector.
func (v Vector4) IsZero() bool {
	panic("unimplemented")
}

// Length implements Vector.
func (v Vector4) Length() float32 {
	panic("unimplemented")
}

// Lerp implements Vector.
func (v Vector4) Lerp(v2 Vector, t float32) Vector {
	panic("unimplemented")
}

// Multiply implements Vector.
func (v Vector4) Multiply(v2 Vector) Vector {
	panic("unimplemented")
}

// MultiplyScalar implements Vector.
func (v Vector4) MultiplyScalar(s float32) Vector {
	panic("unimplemented")
}

// Negate implements Vector.
func (v Vector4) Negate() Vector {
	panic("unimplemented")
}

// Normalize implements Vector.
func (v Vector4) Normalize() Vector {
	panic("unimplemented")
}

// Rotate implements Vector.
func (v Vector4) Rotate(yaw float32, pitch float32, roll float32) Vector {
	panic("unimplemented")
}

// Scale implements Vector.
func (v Vector4) Scale(s float32) Vector {
	panic("unimplemented")
}

// String implements Vector.
func (v Vector4) String() string {
	panic("unimplemented")
}

// Sub implements Vector.
func (v Vector4) Sub(v2 Vector) Vector {
	panic("unimplemented")
}

var _ Vector = Vector4{}
