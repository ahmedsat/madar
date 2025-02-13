package madar

type Matrix3X3 [9]float32

var _ Matrix = Matrix3X3{}

// Ptr implements Matrix.
func (m Matrix3X3) GL() *float32 {
	panic("unimplemented")
}

// MultiplyVector implements Matrix.
func (m Matrix3X3) MultiplyVector(v Vector) Vector {
	panic("unimplemented")
}

func IdentityMatrix3X3() Matrix3X3 {
	return Matrix3X3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

// Add implements Matrix.
func (m Matrix3X3) Add(Matrix) Matrix {
	panic("unimplemented")
}

// Determinant implements Matrix.
func (m Matrix3X3) Determinant() float32 {
	panic("unimplemented")
}

// DivideScalar implements Matrix.
func (m Matrix3X3) DivideScalar(s float32) Matrix {
	panic("unimplemented")
}

// Inverse implements Matrix.
func (m Matrix3X3) Inverse() (Matrix, bool) {
	panic("unimplemented")
}

// IsCholesky implements Matrix.
func (m Matrix3X3) IsCholesky() bool {
	panic("unimplemented")
}

// IsEqual implements Matrix.
func (m Matrix3X3) IsEqual(Matrix) bool {
	panic("unimplemented")
}

// IsHessenberg implements Matrix.
func (m Matrix3X3) IsHessenberg() bool {
	panic("unimplemented")
}

// IsIdentity implements Matrix.
func (m Matrix3X3) IsIdentity() bool {
	panic("unimplemented")
}

// IsLowerTriangular implements Matrix.
func (m Matrix3X3) IsLowerTriangular() bool {
	panic("unimplemented")
}

// IsOrthogonal implements Matrix.
func (m Matrix3X3) IsOrthogonal() bool {
	panic("unimplemented")
}

// IsSymmetric implements Matrix.
func (m Matrix3X3) IsSymmetric() bool {
	panic("unimplemented")
}

// IsTriangular implements Matrix.
func (m Matrix3X3) IsTriangular() bool {
	panic("unimplemented")
}

// IsUpperTriangular implements Matrix.
func (m Matrix3X3) IsUpperTriangular() bool {
	panic("unimplemented")
}

// IsZero implements Matrix.
func (m Matrix3X3) IsZero() bool {
	panic("unimplemented")
}

// Multiply implements Matrix.
func (m Matrix3X3) Multiply(Matrix) Matrix {
	panic("unimplemented")
}

// MultiplyScalar implements Matrix.
func (m Matrix3X3) MultiplyScalar(s float32) Matrix {
	panic("unimplemented")
}

// Negate implements Matrix.
func (m Matrix3X3) Negate() Matrix {
	panic("unimplemented")
}

// Sub implements Matrix.
func (m Matrix3X3) Sub(Matrix) Matrix {
	panic("unimplemented")
}

// Trace implements Matrix.
func (m Matrix3X3) Trace() float32 {
	panic("unimplemented")
}

// Transpose implements Matrix.
func (m Matrix3X3) Transpose() Matrix {
	panic("unimplemented")
}

// String implements fmt.Stringer.
func (m Matrix3X3) String() string {
	panic("unimplemented")
}
