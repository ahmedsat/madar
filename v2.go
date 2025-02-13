package madar

type Vector2 struct {
	X, Y float32
}

// Add implements Vector.
func (v Vector2) Add(v2 Vector) Vector {
	panic("unimplemented")
}

// Distance implements Vector.
func (v Vector2) Distance(v2 Vector) float32 {
	panic("unimplemented")
}

// Divide implements Vector.
func (v Vector2) Divide(v2 Vector) Vector {
	panic("unimplemented")
}

// DivideScalar implements Vector.
func (v Vector2) DivideScalar(s float32) Vector {
	panic("unimplemented")
}

// Dot implements Vector.
func (v Vector2) Dot(v2 Vector) float32 {
	panic("unimplemented")
}

// IsEqual implements Vector.
func (v Vector2) IsEqual(v2 Vector) bool {
	panic("unimplemented")
}

// IsZero implements Vector.
func (v Vector2) IsZero() bool {
	panic("unimplemented")
}

// Length implements Vector.
func (v Vector2) Length() float32 {
	panic("unimplemented")
}

// Lerp implements Vector.
func (v Vector2) Lerp(v2 Vector, t float32) Vector {
	panic("unimplemented")
}

// Multiply implements Vector.
func (v Vector2) Multiply(v2 Vector) Vector {
	panic("unimplemented")
}

// MultiplyScalar implements Vector.
func (v Vector2) MultiplyScalar(s float32) Vector {
	panic("unimplemented")
}

// Negate implements Vector.
func (v Vector2) Negate() Vector {
	panic("unimplemented")
}

// Normalize implements Vector.
func (v Vector2) Normalize() Vector {
	panic("unimplemented")
}

// Rotate implements Vector.
func (v Vector2) Rotate(yaw float32, pitch float32, roll float32) Vector {
	panic("unimplemented")
}

// Scale implements Vector.
func (v Vector2) Scale(s float32) Vector {
	panic("unimplemented")
}

// String implements Vector.
func (v Vector2) String() string {
	panic("unimplemented")
}

// Sub implements Vector.
func (v Vector2) Sub(v2 Vector) Vector {
	panic("unimplemented")
}

var _ Vector = Vector2{}
