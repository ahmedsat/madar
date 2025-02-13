package madar

import "fmt"

type Matrix4X4 [16]float32

var _ Matrix = Matrix4X4{}

// Ptr implements Matrix.
func (m Matrix4X4) GL() *float32 {
	transposed := m.Transpose().(Matrix4X4)
	return &transposed[0]
}

// MultiplyVector implements Matrix.
func (m Matrix4X4) MultiplyVector(v Vector) Vector {
	panic("unimplemented")
}

func IdentityMatrix4X4() Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Add implements Matrix.
func (m Matrix4X4) Add(Matrix) Matrix {
	panic("unimplemented")
}

// Determinant implements Matrix.
func (m Matrix4X4) Determinant() float32 {
	panic("unimplemented")
}

// DivideScalar implements Matrix.
func (m Matrix4X4) DivideScalar(s float32) Matrix {
	panic("unimplemented")
}

// Inverse implements Matrix.
func (m Matrix4X4) Inverse() (Matrix, bool) {
	panic("unimplemented")
}

// IsCholesky implements Matrix.
func (m Matrix4X4) IsCholesky() bool {
	panic("unimplemented")
}

// IsEqual implements Matrix.
func (m Matrix4X4) IsEqual(Matrix) bool {
	panic("unimplemented")
}

// IsHessenberg implements Matrix.
func (m Matrix4X4) IsHessenberg() bool {
	panic("unimplemented")
}

// IsIdentity implements Matrix.
func (m Matrix4X4) IsIdentity() bool {
	panic("unimplemented")
}

// IsLowerTriangular implements Matrix.
func (m Matrix4X4) IsLowerTriangular() bool {
	panic("unimplemented")
}

// IsOrthogonal implements Matrix.
func (m Matrix4X4) IsOrthogonal() bool {
	panic("unimplemented")
}

// IsSymmetric implements Matrix.
func (m Matrix4X4) IsSymmetric() bool {
	panic("unimplemented")
}

// IsTriangular implements Matrix.
func (m Matrix4X4) IsTriangular() bool {
	panic("unimplemented")
}

// IsUpperTriangular implements Matrix.
func (m Matrix4X4) IsUpperTriangular() bool {
	panic("unimplemented")
}

// IsZero implements Matrix.
func (m Matrix4X4) IsZero() bool {
	panic("unimplemented")
}

// Multiply implements Matrix.
func (m0 Matrix4X4) Multiply(m Matrix) Matrix {
	m1, ok := m.(Matrix4X4)
	Assert(ok, fmt.Sprintf("expected Matrix4X4, got %T", m))

	return Matrix4X4{
		m0[0]*m1[0] + m0[4]*m1[1] + m0[8]*m1[2] + m0[12]*m1[3],
		m0[1]*m1[0] + m0[5]*m1[1] + m0[9]*m1[2] + m0[13]*m1[3],
		m0[2]*m1[0] + m0[6]*m1[1] + m0[10]*m1[2] + m0[14]*m1[3],
		m0[3]*m1[0] + m0[7]*m1[1] + m0[11]*m1[2] + m0[15]*m1[3],

		m0[0]*m1[4] + m0[4]*m1[5] + m0[8]*m1[6] + m0[12]*m1[7],
		m0[1]*m1[4] + m0[5]*m1[5] + m0[9]*m1[6] + m0[13]*m1[7],
		m0[2]*m1[4] + m0[6]*m1[5] + m0[10]*m1[6] + m0[14]*m1[7],
		m0[3]*m1[4] + m0[7]*m1[5] + m0[11]*m1[6] + m0[15]*m1[7],

		m0[0]*m1[8] + m0[4]*m1[9] + m0[8]*m1[10] + m0[12]*m1[11],
		m0[1]*m1[8] + m0[5]*m1[9] + m0[9]*m1[10] + m0[13]*m1[11],
		m0[2]*m1[8] + m0[6]*m1[9] + m0[10]*m1[10] + m0[14]*m1[11],
		m0[3]*m1[8] + m0[7]*m1[9] + m0[11]*m1[10] + m0[15]*m1[11],

		m0[0]*m1[12] + m0[4]*m1[13] + m0[8]*m1[14] + m0[12]*m1[15],
		m0[1]*m1[12] + m0[5]*m1[13] + m0[9]*m1[14] + m0[13]*m1[15],
		m0[2]*m1[12] + m0[6]*m1[13] + m0[10]*m1[14] + m0[14]*m1[15],
		m0[3]*m1[12] + m0[7]*m1[13] + m0[11]*m1[14] + m0[15]*m1[15],
	}
}

// MultiplyScalar implements Matrix.
func (m Matrix4X4) MultiplyScalar(s float32) Matrix {
	panic("unimplemented")
}

// Negate implements Matrix.
func (m Matrix4X4) Negate() Matrix {
	panic("unimplemented")
}

// String implements Matrix.
func (m Matrix4X4) String() string {
	panic("unimplemented")
}

// Sub implements Matrix.
func (m Matrix4X4) Sub(Matrix) Matrix {
	panic("unimplemented")
}

// Trace implements Matrix.
func (m Matrix4X4) Trace() float32 {
	panic("unimplemented")
}

// Transpose implements Matrix.
func (m Matrix4X4) Transpose() Matrix {
	return Matrix4X4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15],
	}
}
