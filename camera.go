package main

import (
	"fmt"
	"image"
	"math"
)

type camera struct {
	aspect_ratio              float64
	image_width, image_height uint
	center, pixel00_loc       Point3
	pixel_du, pixel_dv        Vec3
	samples_per_pixel         uint
	pixel_samples_scale       float64
	max_depth                 uint
	vfov                      float64
	lookfrom, lookat          Point3
	vup, u, v, w              Vec3
	defocus_angle, focus_dist float64
	defocus_disk_u            Vec3
	defocus_disk_v            Vec3
}

func NewCamera(aspect_ratio float64, image_width uint, samples_per_pixel, max_depth uint, vfov float64, lookfrom, lookat Point3, vup Vec3, focus_dist, defocus_angle float64) camera {
	// Calculate image_height automatically (must be >=1)
	image_height := max(1, uint(float64(image_width)/aspect_ratio))

	// Viewport widths less than one are ok since they are real valued.
	center := lookfrom

	theta := DegToRad(vfov)
	h := math.Tan(theta / 2)
	viewport_h := 2 * h * focus_dist
	viewport_w := viewport_h * float64(image_width) / float64(image_height)

	w := lookfrom.Sub(lookat).UnitVec()
	u := vup.Cross(w).UnitVec()
	v := w.Cross(u)

	// Define the viewport as vectors
	viewport_u := u.Scale(viewport_w)
	viewport_v := v.Scale(-viewport_h)

	pixel_du := viewport_u.Scale(1.0 / float64(image_width))
	pixel_dv := viewport_v.Scale(1.0 / float64(image_height))

	viewport_upper_left := center.Sub(w.Scale(focus_dist), viewport_u.Scale(0.5), viewport_v.Scale(0.5))
	pixel00_loc := viewport_upper_left.Add(pixel_du.Add(pixel_dv).Scale(0.5))
	defocus_radius := focus_dist * math.Tan(DegToRad(defocus_angle / 2.0))
	defocus_disk_u := u.Scale(defocus_radius)
	defocus_disk_v := v.Scale(defocus_radius)

	pixel_samples_scale := 1.0 / float64(samples_per_pixel)
	return camera{aspect_ratio, image_width, uint(image_height), center, pixel00_loc, pixel_du, pixel_dv, samples_per_pixel, pixel_samples_scale, max_depth, vfov, lookfrom, lookat, vup, u, v, w, defocus_angle, focus_dist, defocus_disk_u, defocus_disk_v}

}

func (c camera) Render(world HittableList) image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(c.image_width), int(c.image_height)))

	for row := range c.image_height {
		for col := range c.image_width {

			var pixel_color Color

			for range c.samples_per_pixel {
				r := c.GetRay(row, col)
				rc := RayColor(r, c.max_depth, world)
				pixel_color = pixel_color.Add(rc)
			}

			pixel_color.Scale(c.pixel_samples_scale).Write(col, row, img)
		}
		fmt.Printf("%d/%d\n", row, c.image_height)
	}

	return *img
}

func (c camera) GetRay(row, col uint) Ray {
	offset := SampleSquare()
	pixel_sample := c.pixel00_loc.Add(c.pixel_du.Scale(float64(col)+offset.x), c.pixel_dv.Scale(float64(row)+offset.y))

	var ray_origin Point3
	if c.defocus_angle <= 0 {
		ray_origin = c.center
	} else {
		ray_origin = c.DefocusDiskSample()
	}
	ray_dir := pixel_sample.Sub(ray_origin)

	return Ray{ray_origin, ray_dir}
}

func (c camera) DefocusDiskSample() Point3 {
	p := RandomVecInUnitDisk()
	return c.center.Add(c.defocus_disk_u.Scale(p.x), c.defocus_disk_v.Scale(p.y))
}

func RayColor(r Ray, depth uint, world Hittable) Color {
	if depth == 0 {
		return Color{0, 0, 0}
	}

	var rec HitRecord

	if world.Hit(r, interval{0.001, math.Inf(1)}, &rec) {
		var scattered Ray
		var attenuation Color

		if rec.mat.scatter(r, rec, &attenuation, &scattered) {
			return attenuation.Mul(RayColor(scattered, depth-1, world))
		}

		return Color{0, 0, 0}
	}

	// Background
	unit_dir := r.dir.UnitVec()
	a := 0.5 * (unit_dir.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.Scale(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.Scale(a))
}
