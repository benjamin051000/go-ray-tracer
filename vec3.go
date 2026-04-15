package main

import "math"

type Vec3 struct {
	x, y, z float64
}

type Point3 Vec3

func Add(l, r Vec3) Vec3 {
	return Vec3{l.x + r.x, l.y + r.y, l.z + r.z}
}

func Sub(l, r Vec3) Vec3 {
	return Vec3{l.x - r.x, l.y - r.y, l.z - r.z}
}

func Mul(l, r Vec3) Vec3 {
	return Vec3{l.x * r.x, l.y * r.y, l.z * r.z}
}

func Div(l, r Vec3) Vec3 {
	return Vec3{l.x / r.x, l.y / r.y, l.z / r.z}
}

func Scale(v Vec3, t float64) Vec3 {
	return Vec3{v.x * t, v.y * t, v.z * t}
}

func (v Vec3) LenSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

// Whoa, this is the same as LenSquared...
func Dot(l, r Vec3) float64 {
	return l.x*r.x + l.y*r.y + l.z*r.z
}

func Cross(u, v Vec3) Vec3 {
	return Vec3{u.y*v.z - u.z*v.y,
		u.z*v.x - u.x*v.z,
		u.x*v.y - u.y*v.x}
}

func (v Vec3) UnitVec() Vec3 {
	l := v.Len()
	return Div(v, Vec3{l, l, l})
}
