package main

type material interface {
	scatter(r_in Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool
}

type lambertian struct {
	albedo Color
}

func (l lambertian) scatter(r_in Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	scatter_direction := rec.normal.Add(RandomUnitVec3())
	if scatter_direction.NearZero() {
		scatter_direction = rec.normal
	}

	// TODO should this just be a retval?
	*scattered = Ray{rec.p, scatter_direction}
	*attenuation = l.albedo
	return true
}

type metal struct {
	albedo Color
	fuzz float64
}

func (m metal) scatter(r_in Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := r_in.dir.Reflect(rec.normal)
	reflected = reflected.UnitVec().Add(RandomUnitVec3().Scale(m.fuzz))
	*scattered = Ray{rec.p, reflected}
	*attenuation = m.albedo
	return scattered.dir.Dot(rec.normal) > 0
}
