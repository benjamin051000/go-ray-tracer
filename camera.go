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
}

func NewCamera(aspect_ratio float64, image_width uint) camera {
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

	return camera{aspect_ratio, image_width, uint(image_height), camera_center, pixel00_loc, pixel_du, pixel_dv}

}

func (c camera) Render(world HittableList) image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, int(c.image_width), int(c.image_height)))

	for row := range c.image_height {
		for col := range c.image_width {
			pixel_center := c.pixel00_loc.Add(c.pixel_du.Scale(float64(col)), c.pixel_dv.Scale(float64(row)))
			ray_direction := pixel_center.Sub(Vec3(c.center)) // technically doesn't need to be a unit vector

			r := Ray{c.center, ray_direction}
			pixel_color := RayColor(r, world)

			pixel_color.Write(col, row, img)
		}
		fmt.Printf("%d/%d\n", row, c.image_height)
	}

	return *img
}

func RayColor(r Ray, world Hittable) Color {
	var rec HitRecord
	if world.Hit(r, interval{0, math.Inf(1)}, &rec) {
		return rec.normal.Add(Vec3{1, 1, 1}).Scale(0.5)
	}

	// Background
	unit_dir := r.dir.UnitVec()
	a := 0.5 * (unit_dir.y + 1.0)
	return Vec3{1.0, 1.0, 1.0}.Scale(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.Scale(a))
}
