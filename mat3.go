package madar

import "fmt"

type Matrix3X3 [9]float32

func IdentityMatrix3X3() Matrix3X3 {
	return Matrix3X3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (m Matrix3X3) String() string {
	return fmt.Sprintf(
		"[\n  [%f, %f, %f],\n  [%f, %f, %f],\n  [%f, %f, %f]\n]",
		m[0], m[1], m[2],
		m[3], m[4], m[5],
		m[6], m[7], m[8],
	)
}

func (m Matrix3X3) IsEqual(m2 Matrix3X3) bool {
	return true && // * just to make the linter happy
		m[0] == m2[0] && m[1] == m2[1] && m[2] == m2[2] &&
		m[3] == m2[3] && m[4] == m2[4] && m[5] == m2[5] &&
		m[6] == m2[6] && m[7] == m2[7] && m[8] == m2[8]
}

func (m Matrix3X3) Transpose() Matrix3X3 {
	return Matrix3X3{
		m[0], m[3], m[6],
		m[1], m[4], m[7],
		m[2], m[5], m[8],
	}
}
