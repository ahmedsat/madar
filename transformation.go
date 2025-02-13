package madar

import "fmt"

func ScalingMatrix(sx, sy, sz float32) Matrix4X4 {
	return Matrix4X4{
		sx, 0, 0, 0,
		0, sy, 0, 0,
		0, 0, sz, 0,
		0, 0, 0, 1,
	}
}

func TranslationMatrix(tx, ty, tz float32) Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, tx,
		0, 1, 0, ty,
		0, 0, 1, tz,
		0, 0, 0, 1,
	}
}

func RotationMatrixX(deg float32) Matrix4X4 {
	return Matrix4X4{
		1, 0, 0, 0,
		0, Cos(deg), -Sin(deg), 0,
		0, Sin(deg), Cos(deg), 0,
		0, 0, 0, 1,
	}
}

func RotationMatrixY(deg float32) Matrix4X4 {
	return Matrix4X4{
		Cos(deg), 0, Sin(deg), 0,
		0, 1, 0, 0,
		-Sin(deg), 0, Cos(deg), 0,
		0, 0, 0, 1,
	}
}

func RotationMatrixZ(deg float32) Matrix4X4 {
	return Matrix4X4{
		Cos(deg), -Sin(deg), 0, 0,
		Sin(deg), Cos(deg), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotationMatrix(degX, degY, degZ float32) Matrix4X4 {

	m := RotationMatrixX(degX).
		Multiply(RotationMatrixY(degY)).
		Multiply(RotationMatrixZ(degZ))

	m4, ok := m.(Matrix4X4)
	Assert(ok, fmt.Sprintf("expected Matrix4X4, got %T", m))

	return m4
}
