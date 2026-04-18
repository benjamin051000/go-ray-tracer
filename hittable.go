package main

type HitRecord struct {
	p      Point3
	normal Vec3
	t      float64
}

type Hittable interface {
	Hit(r Ray, r_tmin, r_tmax float64, rec *HitRecord) bool
}
