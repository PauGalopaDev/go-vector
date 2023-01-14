package pgvector

import (
	"fmt"
	"math"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type Vec2[T Numeric] [2]T

func (v *Vec2[T]) Set(x, y T) *Vec2[T] {
	v[0], v[1] = x, y
	return v
}

func (v *Vec2[T]) Get() Vec2[T] {
	return Vec2[T]{v[0], v[1]}
}

func (v *Vec2[T]) Add(v2 Vec2[T]) *Vec2[T] {
	v[0] += v2[0]
	v[1] += v2[1]
	return v
}

func (v *Vec2[T]) Adds(s T) *Vec2[T] {
	v[0] += s
	v[1] += s
	return v
}

func (v *Vec2[T]) Sub(v2 Vec2[T]) *Vec2[T] {
	v[0] -= v2[0]
	v[1] -= v2[1]
	return v
}

func (v *Vec2[T]) Subs(s T) *Vec2[T] {
	v[0] -= s
	v[1] -= s
	return v
}

func (v *Vec2[T]) Hadamard(v2 Vec2[T]) *Vec2[T] {
	v[0] *= v2[0]
	v[1] *= v2[1]
	return v
}

func (v *Vec2[T]) Muls(s T) *Vec2[T] {
	v[0] *= s
	v[1] *= s
	return v
}

func (v *Vec2[T]) Divs(s T) *Vec2[T] {
	v[0] /= s
	v[1] /= s
	return v
}

func (v *Vec2[T]) Rotate(a float64) *Vec2[T] {
	cos, sin := math.Cos(a), math.Sin(a)
	v[0], v[1] = T(cos*float64(v[0])-sin*float64(v[1])), T(sin*float64(v[0])+cos*float64(v[1]))
	return v
}

func (v *Vec2[T]) Unit() *Vec2[T] {
	norm := T(v.Norm())
	v[0] /= norm
	v[1] /= norm
	return v
}

func (v *Vec2[T]) Dot(v2 Vec2[T]) float64 {
	return float64(v[0]*v2[0] + v[1]*v2[1])
}

func (v *Vec2[T]) Norm() float64 {
	return math.Sqrt(float64(v[0]*v[0] + v[1]*v[1]))
}

func (v *Vec2[T]) Angle(v2 Vec2[T]) float64 {
	return math.Acos(v.Dot(v2) / (v.Norm() * v2.Norm()))

}

func (v *Vec2[T]) String() string {
	return fmt.Sprintf("Vec2(%v, %v)", v[0], v[1])
}

func (v *Vec2[T]) All(f func(T) bool) bool {
	return f(v[0]) && f(v[1])
}

func (v *Vec2[T]) Any(f func(T) bool) bool {
	return f(v[0]) || f(v[1])
}

func (v *Vec2[T]) X() *T {
	return &v[0]
}

func (v *Vec2[T]) Y() *T {
	return &v[1]
}

func (v *Vec2[T]) XY() (T, T) {
	return v[0], v[1]
}

func (v *Vec2[T]) YX() (T, T) {
	return v[1], v[0]
}
