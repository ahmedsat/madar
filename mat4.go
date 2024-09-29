package madar

import (
	"math"
)

// Matrix4 represents a 4x4 matrix in row-major order.
type Matrix4 [16]float32

// NewMatrix4 creates a new Matrix4 with the given values.
func NewMatrix4(values [16]float32) Matrix4 {
	return Matrix4(values)
}

// Identity returns an identity matrix.
func Identity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Multiply multiplies two matrices and returns the result.
func (m Matrix4) Multiply(n Matrix4) Matrix4 {
	var result Matrix4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := float32(0)
			for i := 0; i < 4; i++ {
				sum += m[row*4+i] * n[i*4+col]
			}
			result[row*4+col] = sum
		}
	}
	return result
}

// MultiplyVector multiplies the matrix with a Vector4.
func (m Matrix4) MultiplyVector(v [4]float32) [4]float32 {
	var result [4]float32
	for row := 0; row < 4; row++ {
		sum := float32(0)
		for col := 0; col < 4; col++ {
			sum += m[row*4+col] * v[col]
		}
		result[row] = sum
	}
	return result
}

// Transpose returns the transpose of the matrix.
func (m Matrix4) Transpose() Matrix4 {
	var result Matrix4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			result[col*4+row] = m[row*4+col]
		}
	}
	return result
}

// Determinant calculates the determinant of the matrix.
func (m Matrix4) Determinant() float32 {
	return m[0]*m[5]*m[10]*m[15] +
		m[0]*m[9]*m[14]*m[7] +
		m[0]*m[13]*m[6]*m[11] -
		m[0]*m[13]*m[10]*m[7] -
		m[0]*m[5]*m[14]*m[11] -
		m[0]*m[9]*m[6]*m[15] -
		m[4]*m[1]*m[10]*m[15] -
		m[8]*m[1]*m[14]*m[7] -
		m[12]*m[1]*m[6]*m[11] +
		m[12]*m[1]*m[10]*m[7] +
		m[4]*m[1]*m[14]*m[11] +
		m[8]*m[1]*m[6]*m[15] +
		m[4]*m[9]*m[2]*m[15] +
		m[8]*m[13]*m[2]*m[7] +
		m[12]*m[5]*m[2]*m[11] -
		m[12]*m[9]*m[2]*m[7] -
		m[4]*m[13]*m[2]*m[11] -
		m[8]*m[5]*m[2]*m[15] -
		m[4]*m[9]*m[14]*m[3] -
		m[8]*m[13]*m[6]*m[3] -
		m[12]*m[5]*m[10]*m[3] +
		m[12]*m[9]*m[6]*m[3] +
		m[4]*m[13]*m[10]*m[3] +
		m[8]*m[5]*m[14]*m[3]
}

// Inverse returns the inverse of the matrix and a boolean indicating success.
func (m Matrix4) Inverse() (Matrix4, bool) {
	det := m.Determinant()
	if math.Abs(float64(det)) < 1e-6 {
		return Matrix4{}, false
	}

	invDet := 1.0 / det
	result := Matrix4{
		(m[5]*m[10]*m[15] + m[9]*m[14]*m[7] + m[13]*m[6]*m[11] - m[13]*m[10]*m[7] - m[5]*m[14]*m[11] - m[9]*m[6]*m[15]) * invDet,
		(m[1]*m[14]*m[11] + m[9]*m[2]*m[15] + m[13]*m[10]*m[3] - m[13]*m[2]*m[11] - m[1]*m[10]*m[15] - m[9]*m[14]*m[3]) * invDet,
		(m[1]*m[6]*m[15] + m[5]*m[14]*m[3] + m[13]*m[2]*m[7] - m[13]*m[6]*m[3] - m[1]*m[14]*m[7] - m[5]*m[2]*m[15]) * invDet,
		(m[1]*m[10]*m[7] + m[5]*m[2]*m[11] + m[9]*m[6]*m[3] - m[9]*m[2]*m[7] - m[1]*m[6]*m[11] - m[5]*m[10]*m[3]) * invDet,
		(m[4]*m[14]*m[11] + m[8]*m[6]*m[15] + m[12]*m[10]*m[7] - m[12]*m[6]*m[11] - m[4]*m[10]*m[15] - m[8]*m[14]*m[7]) * invDet,
		(m[0]*m[10]*m[15] + m[8]*m[14]*m[3] + m[12]*m[2]*m[11] - m[12]*m[10]*m[3] - m[0]*m[14]*m[11] - m[8]*m[2]*m[15]) * invDet,
		(m[0]*m[14]*m[7] + m[4]*m[2]*m[15] + m[12]*m[6]*m[3] - m[12]*m[2]*m[7] - m[0]*m[6]*m[15] - m[4]*m[14]*m[3]) * invDet,
		(m[0]*m[6]*m[11] + m[4]*m[10]*m[3] + m[8]*m[2]*m[7] - m[8]*m[6]*m[3] - m[0]*m[10]*m[7] - m[4]*m[2]*m[11]) * invDet,
		(m[4]*m[9]*m[15] + m[8]*m[13]*m[7] + m[12]*m[5]*m[11] - m[12]*m[9]*m[7] - m[4]*m[13]*m[11] - m[8]*m[5]*m[15]) * invDet,
		(m[0]*m[13]*m[11] + m[8]*m[1]*m[15] + m[12]*m[9]*m[3] - m[12]*m[1]*m[11] - m[0]*m[9]*m[15] - m[8]*m[13]*m[3]) * invDet,
		(m[0]*m[5]*m[15] + m[4]*m[13]*m[3] + m[12]*m[1]*m[7] - m[12]*m[5]*m[3] - m[0]*m[13]*m[7] - m[4]*m[1]*m[15]) * invDet,
		(m[0]*m[9]*m[7] + m[4]*m[1]*m[11] + m[8]*m[5]*m[3] - m[8]*m[1]*m[7] - m[0]*m[5]*m[11] - m[4]*m[9]*m[3]) * invDet,
		(m[4]*m[13]*m[10] + m[8]*m[5]*m[14] + m[12]*m[9]*m[6] - m[12]*m[5]*m[10] - m[4]*m[9]*m[14] - m[8]*m[13]*m[6]) * invDet,
		(m[0]*m[9]*m[14] + m[8]*m[13]*m[2] + m[12]*m[1]*m[10] - m[12]*m[9]*m[2] - m[0]*m[13]*m[10] - m[8]*m[1]*m[14]) * invDet,
		(m[0]*m[13]*m[6] + m[4]*m[1]*m[14] + m[12]*m[5]*m[2] - m[12]*m[1]*m[6] - m[0]*m[5]*m[14] - m[4]*m[13]*m[2]) * invDet,
		(m[0]*m[5]*m[10] + m[4]*m[9]*m[2] + m[8]*m[1]*m[6] - m[8]*m[5]*m[2] - m[0]*m[9]*m[6] - m[4]*m[1]*m[10]) * invDet,
	}
	return result, true
}

// Translation returns a translation matrix.
func Translation(x, y, z float32) Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		x, y, z, 1,
	}
}

// Scale returns a scaling matrix.
func Scale(x, y, z float32) Matrix4 {
	return Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

// RotationX returns a rotation matrix around the X-axis.
func RotationX(angle float32) Matrix4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

// RotationY returns a rotation matrix around the Y-axis.
func RotationY(angle float32) Matrix4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

// RotationZ returns a rotation matrix around the Z-axis.
func RotationZ(angle float32) Matrix4 {
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	return Matrix4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Perspective returns a perspective projection matrix.
func Perspective(fov, aspect, near, far float32) Matrix4 {
	tanHalfFov := float32(math.Tan(float64(fov) / 2))
	return Matrix4{
		1 / (aspect * tanHalfFov), 0, 0, 0,
		0, 1 / tanHalfFov, 0, 0,
		0, 0, -(far + near) / (far - near), -1,
		0, 0, -(2 * far * near) / (far - near), 0,
	}
}

// Orthographic returns an orthographic projection matrix.
func Orthographic(left, right, bottom, top, near, far float32) Matrix4 {
	return Matrix4{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, -2 / (far - near), 0,
		-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1,
	}
}

// LookAt returns a view matrix.
func LookAt(eye, target, up [3]float32) Matrix4 {
	zAxis := normalize(subtract(eye, target))
	xAxis := normalize(cross(up, zAxis))
	yAxis := cross(zAxis, xAxis)

	return Matrix4{
		xAxis[0], yAxis[0], zAxis[0], 0,
		xAxis[1], yAxis[1], zAxis[1], 0,
		xAxis[2], yAxis[2], zAxis[2], 0,
		-dot(xAxis, eye), -dot(yAxis, eye), -dot(zAxis, eye), 1,
	}
}

// Helper functions for vector operations
func subtract(a, b [3]float32) [3]float32 {
	return [3]float32{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func normalize(v [3]float32) [3]float32 {
	length := float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
	return [3]float32{v[0] / length, v[1] / length, v[2] / length}
}

func cross(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func dot(a, b [3]float32) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// OrthographicMatrix creates an orthographic projection matrix.
// Parameters:
// - left, right: specify the coordinates for the left and right vertical clipping planes
// - bottom, top: specify the coordinates for the bottom and top horizontal clipping planes
// - near, far: specify the distances to the nearer and farther depth clipping planes
func OrthographicMatrix(left, right, bottom, top, near, far float32) Matrix4 {
	// Check for division by zero
	if right == left || top == bottom || far == near {
		panic("Invalid parameters for orthographic matrix: division by zero")
	}

	return Matrix4{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, -2 / (far - near), 0,
		-(right + left) / (right - left),
		-(top + bottom) / (top - bottom),
		-(far + near) / (far - near),
		1,
	}
}

// PerspectiveMatrix creates a perspective projection matrix.
// Parameters:
// - fov: Field of view in radians
// - aspect: Aspect ratio (width / height) of the view plane
// - near: Distance to the near clipping plane (must be positive)
// - far: Distance to the far clipping plane (must be greater than near)
func PerspectiveMatrix(fov, aspect, near, far float32) Matrix4 {
	// Check for invalid parameters
	if near <= 0 || far <= near {
		panic("Invalid parameters for perspective matrix: near must be positive and far must be greater than near")
	}
	if fov <= 0 || fov >= math.Pi {
		panic("Invalid field of view for perspective matrix: must be between 0 and π")
	}
	if aspect <= 0 {
		panic("Invalid aspect ratio for perspective matrix: must be positive")
	}

	tanHalfFov := float32(math.Tan(float64(fov) / 2))
	f := 1.0 / tanHalfFov
	d := 1.0 / (near - far)

	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) * d, -1,
		0, 0, 2 * far * near * d, 0,
	}
}
