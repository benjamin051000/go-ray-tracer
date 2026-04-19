package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z float64
}

type Point3 = Vec3

func SampleSquare() Vec3 {
	return Vec3{rand.Float64() - 0.5, rand.Float64() + 0.5, 0}
}

func RandomVec3() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func RandomVec3Range(min, max float64) Vec3 {
	return Vec3{RandFloatRange(min, max), RandFloatRange(min, max), RandFloatRange(min, max)}
}

func RandomUnitVec3() Vec3 {
	for {
		p := RandomVec3Range(-1, 1)
		lensq := p.LenSquared()
		if 1e-160 < lensq && lensq <= 1 {
			return p.Scale(1 / math.Sqrt(lensq))
		}
	}
}

func RandomVecOnHemisphere(normal Vec3) Vec3 {
	on_unit_sphere := RandomUnitVec3()
	if on_unit_sphere.Dot(normal) > 0 {
		// In the same hemisphere as the normal.
		return on_unit_sphere
	}

	// Not so. Flip it so it is.
	return on_unit_sphere.Scale(-1)
}

func (v Vec3) Add(vecs ...Vec3) Vec3 {
	for _, vec := range vecs {
		v.x += vec.x
		v.y += vec.y
		v.z += vec.z
	}

	return v
}

func (v Vec3) Sub(vecs ...Vec3) Vec3 {
	for _, vec := range vecs {
		v.x -= vec.x
		v.y -= vec.y
		v.z -= vec.z
	}

	return v
}

func (v Vec3) Mul(vecs ...Vec3) Vec3 {
	for _, vec := range vecs {
		v.x *= vec.x
		v.y *= vec.y
		v.z *= vec.z
	}

	return v
}

func (v Vec3) Div(vecs ...Vec3) Vec3 {
	for _, vec := range vecs {
		v.x /= vec.x
		v.y /= vec.y
		v.z /= vec.z
	}

	return v
}

func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.x * t, v.y * t, v.z * t}
}

func (v Vec3) LenSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

// Whoa, this is the same as LenSquared...
func (v Vec3) Dot(vec Vec3) float64 {
	return v.x*vec.x + v.y*vec.y + v.z*vec.z
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{v.y*other.z - v.z*other.y,
		v.z*other.x - v.x*other.z,
		v.x*other.y - v.y*other.x}
}

func (v Vec3) UnitVec() Vec3 {
	l := v.Len()
	return v.Div(Vec3{l, l, l})
}
