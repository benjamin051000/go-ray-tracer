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
	fuzz   float64
}

func (m metal) scatter(r_in Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	reflected := r_in.dir.Reflect(rec.normal)
	reflected = reflected.UnitVec().Add(RandomUnitVec3().Scale(m.fuzz))
	*scattered = Ray{rec.p, reflected}
	*attenuation = m.albedo
	return scattered.dir.Dot(rec.normal) > 0
}

type dielectric struct {
	refraction_index float64
}

func (d dielectric) scatter(r_in Ray, rec HitRecord, attenuation *Color, scattered *Ray) bool {
	*attenuation = Color{1, 1, 1}
	var ri float64
	if rec.front_face {
		ri = 1.0 / d.refraction_index
	} else {
		ri = d.refraction_index
	}

	unit_dir := r_in.dir.UnitVec()
	refracted := unit_dir.Refract(rec.normal, ri)

	*scattered = Ray{rec.p, refracted}
	return true
}
