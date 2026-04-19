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
}

func NewCamera(aspect_ratio float64, image_width uint, samples_per_pixel, max_depth uint) camera {
	// Calculate image_height automatically (must be >=1)
	image_height := max(1, uint(float64(image_width)/aspect_ratio))

	// Viewport widths less than one are ok since they are real valued.
	viewport_h := 2.0
	viewport_w := viewport_h * float64(image_width) / float64(image_height)
	focal_len := 1.0
	camera_center := Point3{0, 0, 0}

	// Define the viewport as vectors
	viewport_u := Vec3{viewport_w, 0, 0}
	viewport_v := Vec3{0, -viewport_h, 0}

	pixel_du := viewport_u.Scale(1.0 / float64(image_width))
	pixel_dv := viewport_v.Scale(1.0 / float64(image_height))

	viewport_upper_left := Vec3(camera_center).Sub(Vec3{0, 0, focal_len}, viewport_u.Scale(0.5), viewport_v.Scale(0.5))
	pixel00_loc := viewport_upper_left.Add(pixel_du.Add(pixel_dv).Scale(0.5))

	pixel_samples_scale := 1.0 / float64(samples_per_pixel)
	return camera{aspect_ratio, image_width, uint(image_height), camera_center, pixel00_loc, pixel_du, pixel_dv, samples_per_pixel, pixel_samples_scale, max_depth}

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

	ray_origin := c.center
	ray_dir := pixel_sample.Sub(ray_origin)

	return Ray{ray_origin, ray_dir}
}

func RayColor(r Ray, depth uint, world Hittable) Color {
	if depth == 0 {
		return Color{0, 0, 0}
	}

	var rec HitRecord
	if world.Hit(r, interval{0.001, math.Inf(1)}, &rec) {
		// return rec.normal.Add(Vec3{1, 1, 1}).Scale(0.5)
		direction := rec.normal.Add(RandomUnitVec3())
		return RayColor(Ray{rec.p, direction}, depth-1, world).Scale(0.5)
	}

	// Background
	unit_dir := r.dir.UnitVec()
	a := 0.5 * (unit_dir.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.Scale(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.Scale(a))
}
