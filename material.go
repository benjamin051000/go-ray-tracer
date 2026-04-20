package main

import (
	"math"
	"math/rand"
)

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
	cos_theta := min(unit_dir.Scale(-1).Dot(rec.normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannot_refract := ri*sin_theta > 1.0

	var direction Vec3
	if cannot_refract || reflectance(cos_theta, ri) > rand.Float64() {
		direction = unit_dir.Reflect(rec.normal)
	} else {
		direction = unit_dir.Refract(rec.normal, ri)
	}

	*scattered = Ray{rec.p, direction}
	return true
}

// reflectance calculated via Schlick's approximation.
func reflectance(cosine, refraction_index float64) float64 {
	r0 := (1 - refraction_index) / (1 + refraction_index)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
