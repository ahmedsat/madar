package madar

import (
	"errors"
	"math"
)

// Vector represents a dimensional agnostic vector
type Vector []float32

// NewVector creates a new Vector with the given components
func NewVector(components ...float32) Vector {
	return Vector(components)
}

// Dimension returns the dimension of the Vector
func (v Vector) Dimension() int {
	return len(v)
}

// Add adds two Vectors of the same dimension
func (v Vector) Add(other Vector) (Vector, error) {
	if v.Dimension() != other.Dimension() {
		return nil, errors.New("vectors must have the same dimension")
	}
	result := make(Vector, v.Dimension())
	for i := range v {
		result[i] = v[i] + other[i]
	}
	return result, nil
}

// Subtract subtracts two Vectors of the same dimension
func (v Vector) Subtract(other Vector) (Vector, error) {
	if v.Dimension() != other.Dimension() {
		return nil, errors.New("vectors must have the same dimension")
	}
	result := make(Vector, v.Dimension())
	for i := range v {
		result[i] = v[i] - other[i]
	}
	return result, nil
}

// Scale multiplies the Vector by a scalar
func (v Vector) Scale(scalar float32) Vector {
	result := make(Vector, v.Dimension())
	for i, val := range v {
		result[i] = val * scalar
	}
	return result
}

// DotProduct calculates the dot product of two Vectors
func (v Vector) DotProduct(other Vector) (float32, error) {
	if v.Dimension() != other.Dimension() {
		return 0, errors.New("vectors must have the same dimension")
	}
	sum := float32(0)
	for i := range v {
		sum += v[i] * other[i]
	}
	return sum, nil
}

// Magnitude calculates the magnitude (length) of the Vector
func (v Vector) Magnitude() float32 {
	sumOfSquares := float32(0)
	for _, val := range v {
		sumOfSquares += val * val
	}
	return float32(math.Sqrt(float64(sumOfSquares)))
}

// Normalize returns a unit Vector in the same direction
func (v Vector) Normalize() Vector {
	mag := v.Magnitude()
	if mag == 0 {
		return v
	}
	return v.Scale(1 / mag)
}
