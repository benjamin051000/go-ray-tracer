package main

type HitRecord struct {
	p          Point3
	normal     Vec3
	mat        material
	t          float64
	front_face bool
}

// SetFaceNormal sets the hit record normal vector.
// NOTE: outward_normal is assumed to have unit length.
func (h *HitRecord) SetFaceNormal(r Ray, outward_normal Vec3) {
	h.front_face = r.dir.Dot(outward_normal) < 0

	if h.front_face {
		h.normal = outward_normal
	} else {
		h.normal = outward_normal.Scale(-1)
	}
}

type Hittable interface {
	Hit(r Ray, ray_t interval, rec *HitRecord) bool
}
