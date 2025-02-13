package madar

type Matrix interface {
	// Basic arithmetic operations
	Add(Matrix) Matrix
	Sub(Matrix) Matrix
	Multiply(Matrix) Matrix
	MultiplyScalar(s float32) Matrix
	MultiplyVector(v Vector) Vector
	DivideScalar(s float32) Matrix
	Negate() Matrix

	// Structural transformations
	Transpose() Matrix
	Trace() float32
	Determinant() float32
	Inverse() (Matrix, bool) // Returns (inverse, success)

	// Matrix properties
	IsEqual(Matrix) bool
	IsIdentity() bool
	IsZero() bool
	IsSymmetric() bool
	IsOrthogonal() bool

	// Triangular matrices
	IsTriangular() bool
	IsUpperTriangular() bool
	IsLowerTriangular() bool

	// Hessenberg, Cholesky, and others (can be separate)
	IsHessenberg() bool
	IsCholesky() bool

	String() string

	GL() *float32
}
