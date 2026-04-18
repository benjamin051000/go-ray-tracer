package main

import "math"

type Sphere struct {
	center Point3
	radius float64 // TODO constructor? book sets to max(0, val)
}

func (s Sphere) Hit(r Ray, ray_t interval, rec *HitRecord) bool {
	oc := Vec3(s.center).Sub(Vec3(r.origin))
	a := r.dir.LenSquared()
	h := Vec3(r.dir).Dot(oc)
	c := oc.LenSquared() - s.radius*s.radius
	discriminant := h*h - a*c
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (h - sqrtd) / a
	if root <= ray_t.min || root >= ray_t.max {
		root = (h + sqrtd) / a
		if root <= ray_t.min || root >= ray_t.max {
			return false
		}
	}

	rec.t = root
	rec.p = r.At(rec.t)
	outward_normal := Vec3(rec.p).Sub(Vec3(s.center)).Scale(1 / s.radius)
	rec.SetFaceNormal(r, outward_normal)

	return true
}
