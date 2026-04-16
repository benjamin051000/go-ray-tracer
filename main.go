package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func RayColor(r Ray) Color {
	unit_dir := r.dir.UnitVec()
	a := 0.5 * (unit_dir.y + 1.0)
	return Color(Vec3{1.0, 1.0, 1.0}.Scale(1.0 - a).Add(Vec3{0.5, 0.7, 1.0}.Scale(a)))
}

func main() {
	// Image

	// aspect_ratio is ideal, due to maths/rounding, viewport ratio may be slightly diff. See 4.2
	aspect_ratio := 16.0 / 9.0
	img_w := 400
	// Calculate img_h automatically (must be >=1)
	img_h := max(1, int(float64(img_w)/aspect_ratio))

	// Camera

	// Viewport widths less than one are ok since they are real valued.
	viewport_h := 2.0
	viewport_w := viewport_h * float64(img_w) / float64(img_h)
	focal_len := 1.0
	camera_center := Point3{0, 0, 0}

	// Define the viewport as vectors
	viewport_u := Vec3{viewport_w, 0, 0}
	viewport_v := Vec3{0, -viewport_h, 0}

	pixel_du := viewport_u.Scale(1.0 / float64(img_w))
	pixel_dv := viewport_v.Scale(1.0 / float64(img_h))

	viewport_upper_left := Vec3(camera_center).Sub(Vec3{0, 0, focal_len}, viewport_u.Scale(0.5), viewport_v.Scale(0.5))
	pixel00_loc := viewport_upper_left.Add(pixel_du.Add(pixel_dv).Scale(0.5))

	img := image.NewRGBA(image.Rect(0, 0, img_w, img_h))

	for row := range img_h {
		for col := range img_w {
			pixel_center := pixel00_loc.Add(pixel_du.Scale(float64(col)), pixel_dv.Scale(float64(row)))
			ray_direction := pixel_center.Sub(Vec3(camera_center)) // technically doesn't need to be a unit vector

			r := Ray{camera_center, ray_direction}
			pixel_color := RayColor(r)

			pixel_color.Write(col, row, img)
		}
	}

	file, err := os.Create("out.png")
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Error encoding image: %v\n", err)
	}
}
