package madar

import (
	"fmt"
	"math"
)

type Matrix4X4 [16]float32

func IdentityMatrix4X4() Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func TranslationMatrix4X4(x, y, z float32) Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

func ScaleMatrix4X4(x, y, z float32) Matrix4X4 {
	return Matrix4X4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

func UniformScaleMatrix4X4(scale float32) Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1 / scale,
	}
}

func RotationMatrix4X4(x, y, z float32) Matrix4X4 {
	cx, sx := float32(math.Cos(float64(x))), float32(math.Sin(float64(x)))
	cy, sy := float32(math.Cos(float64(y))), float32(math.Sin(float64(y)))
	cz, sz := float32(math.Cos(float64(z))), float32(math.Sin(float64(z)))

	return Matrix4X4{
		cy * cz, sx*sy*cz - cx*sz, cx*sy*cz + sx*sz, 0,
		cy * sz, sx*sy*sz + cx*cz, cx*sy*sz - sx*cz, 0,
		-sy, sx * cy, cx * cy, 0,
		0, 0, 0, 1,
	}
}

func PerspectiveMatrix4X4(fov, aspect, near, far float32) Matrix4X4 {
	f := 1.0 / float32(math.Tan(float64(fov/2.0)))
	return Matrix4X4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (far + near) / (near - far), (2 * far * near) / (near - far),
		0, 0, -1, 0,
	}
}

func OrthographicMatrix4X4(left, right, bottom, top, near, far float32) Matrix4X4 {
	return Matrix4X4{
		2 / (right - left), 0, 0, 0,
		0, 2 / (top - bottom), 0, 0,
		0, 0, -2 / (far - near), 0,
		-(right + left) / (right - left), -(top + bottom) / (top - bottom), -(far + near) / (far - near), 1,
	}
}

func LookAtMatrix4X4(eye, center, up Vector3) Matrix4X4 {
	f := center.Sub(eye).Normalize()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)

	return Matrix4X4{
		s.X, s.Y, s.Z, -s.Dot(eye),
		u.X, u.Y, u.Z, -u.Dot(eye),
		-f.X, -f.Y, -f.Z, f.Dot(eye),
		0, 0, 0, 1,
	}
}

func (m Matrix4X4) Multiply(m2 Matrix4X4) Matrix4X4 {
	var result Matrix4X4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				result[i*4+j] += m[i*4+k] * m2[k*4+j]
			}
		}
	}
	return result
}

func (m Matrix4X4) MultiplyVector3(v Vector3) Vector3 {

	return Vector3{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3],
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7],
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11],
	}

}

func (m Matrix4X4) MultiplyVector4(v Vector4) Vector4 {

	return Vector4{
		m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*v.W,
		m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*v.W,
		m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*v.W,
		m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*v.W,
	}

}

func (m Matrix4X4) Transpose() Matrix4X4 {
	var result Matrix4X4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i*4+j] = m[j*4+i]
		}
	}
	return result
}

// String formats the Matrix4X4 as a string in a 4x4 grid for readability.
func (m Matrix4X4) String() string {
	return fmt.Sprintf(
		"[\n  [%f, %f, %f, %f],\n  [%f, %f, %f, %f],\n  [%f, %f, %f, %f],\n  [%f, %f, %f, %f]\n]",
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15],
	)
}
