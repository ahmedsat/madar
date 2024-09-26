package madar

import (
	"errors"
	"math"
)

// Matrix4x4 represents a 4x4 matrix using a fixed-size array of 16 float32 values in row-major order
type Matrix4x4 [16]float32

// NewMatrix4x4 creates a new 4x4 identity matrix
func NewMatrix4x4() Matrix4x4 {
	return Matrix4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Set sets the value at the specified row and column
func (m *Matrix4x4) Set(row, col int, value float32) error {
	if row < 0 || row > 3 || col < 0 || col > 3 {
		return errors.New("index out of bounds")
	}
	m[row*4+col] = value
	return nil
}

// Get retrieves the value at the specified row and column
func (m *Matrix4x4) Get(row, col int) (float32, error) {
	if row < 0 || row > 3 || col < 0 || col > 3 {
		return 0, errors.New("index out of bounds")
	}
	return m[row*4+col], nil
}

// Translate applies a translation to the matrix
func (m *Matrix4x4) Translate(x, y, z float32) {
	m[3] += x*m[0] + y*m[1] + z*m[2]
	m[7] += x*m[4] + y*m[5] + z*m[6]
	m[11] += x*m[8] + y*m[9] + z*m[10]
	m[15] += x*m[12] + y*m[13] + z*m[14]
}

// RotateX applies a rotation around the X-axis to the matrix
func (m *Matrix4x4) RotateX(angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	m4 := m[4]
	m5 := m[5]
	m6 := m[6]
	m7 := m[7]
	m8 := m[8]
	m9 := m[9]
	m10 := m[10]
	m11 := m[11]

	m[4] = m4*c + m8*s
	m[5] = m5*c + m9*s
	m[6] = m6*c + m10*s
	m[7] = m7*c + m11*s
	m[8] = m8*c - m4*s
	m[9] = m9*c - m5*s
	m[10] = m10*c - m6*s
	m[11] = m11*c - m7*s
}

// RotateY applies a rotation around the Y-axis to the matrix
func (m *Matrix4x4) RotateY(angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	m0 := m[0]
	m1 := m[1]
	m2 := m[2]
	m3 := m[3]
	m8 := m[8]
	m9 := m[9]
	m10 := m[10]
	m11 := m[11]

	m[0] = m0*c - m8*s
	m[1] = m1*c - m9*s
	m[2] = m2*c - m10*s
	m[3] = m3*c - m11*s
	m[8] = m0*s + m8*c
	m[9] = m1*s + m9*c
	m[10] = m2*s + m10*c
	m[11] = m3*s + m11*c
}

// RotateZ applies a rotation around the Z-axis to the matrix
func (m *Matrix4x4) RotateZ(angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	m0 := m[0]
	m1 := m[1]
	m2 := m[2]
	m3 := m[3]
	m4 := m[4]
	m5 := m[5]
	m6 := m[6]
	m7 := m[7]

	m[0] = m0*c + m4*s
	m[1] = m1*c + m5*s
	m[2] = m2*c + m6*s
	m[3] = m3*c + m7*s
	m[4] = m4*c - m0*s
	m[5] = m5*c - m1*s
	m[6] = m6*c - m2*s
	m[7] = m7*c - m3*s
}

// Multiply performs matrix multiplication
func (m *Matrix4x4) Multiply(other Matrix4x4) Matrix4x4 {
	var result Matrix4x4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := float32(0)
			for i := 0; i < 4; i++ {
				sum += m[row*4+i] * other[i*4+col]
			}
			result[row*4+col] = sum
		}
	}
	return result
}

// TransformVector4 transforms a 4D vector by this matrix
func (m *Matrix4x4) TransformVector4(x, y, z, w float32) (float32, float32, float32, float32) {
	return m[0]*x + m[1]*y + m[2]*z + m[3]*w,
		m[4]*x + m[5]*y + m[6]*z + m[7]*w,
		m[8]*x + m[9]*y + m[10]*z + m[11]*w,
		m[12]*x + m[13]*y + m[14]*z + m[15]*w
}

// Other methods (Scale, Transpose, Determinant) remain the same as in the previous implementation
