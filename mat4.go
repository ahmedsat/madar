package madar

import (
	"fmt"
	"math"
)

// Matrix4 represents a 4x4 matrix in row-major order.
type Matrix4 [16]float32

// NewMatrix4 creates a new identity matrix in row-major order.
func NewMatrix4() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Set sets the values of the matrix in row-major order.
func (m *Matrix4) Set(values [16]float32) {
	for i := 0; i < 16; i++ {
		m[i] = values[i]
	}
}

// Multiply performs matrix multiplication (m * n) and returns the resulting matrix in row-major order.
func (m Matrix4) Multiply(n Matrix4) Matrix4 {
	var result Matrix4

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			result[row*4+col] = m[row*4+0]*n[col+0*4] +
				m[row*4+1]*n[col+1*4] +
				m[row*4+2]*n[col+2*4] +
				m[row*4+3]*n[col+3*4]
		}
	}

	return result
}

// Translate applies a translation transformation by multiplying
// the current matrix with a translation matrix in row-major order.
func (m *Matrix4) Translate(x, y, z float32) {
	translation := Matrix4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
	*m = m.Multiply(translation)
}

// Scale applies a scaling transformation in row-major order.
func (m *Matrix4) Scale(x, y, z float32) {
	m[0] *= x
	m[5] *= y
	m[10] *= z
}

// RotateZ applies a rotation around the Z-axis (in radians) in row-major order.
func (m *Matrix4) RotateZ(angle float32) {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))

	rotation := Matrix4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	*m = m.Multiply(rotation)
}

// Print displays the matrix in row-major format.
func (m Matrix4) Print() {
	fmt.Printf("[%.2f %.2f %.2f %.2f]\n", m[0], m[1], m[2], m[3])
	fmt.Printf("[%.2f %.2f %.2f %.2f]\n", m[4], m[5], m[6], m[7])
	fmt.Printf("[%.2f %.2f %.2f %.2f]\n", m[8], m[9], m[10], m[11])
	fmt.Printf("[%.2f %.2f %.2f %.2f]\n", m[12], m[13], m[14], m[15])
}
