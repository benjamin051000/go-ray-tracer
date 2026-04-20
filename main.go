package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	// Materials
	// mat_ground := lambertian{Color{0.8, 0.8, 0.0}}
	// mat_center := lambertian{Color{0.1, 0.2, 0.5}}
	// mat_left := dielectric{refraction_index: 1.50}
	// mat_bubble := dielectric{refraction_index: 1.00 / 1.50}
	// mat_right := metal{Color{0.8, 0.6, 0.2}, 1.0}
	mat_left := lambertian{Color{0, 0, 1}}
	mat_right := lambertian{Color{1, 0, 0}}

	R := math.Cos(math.Pi/4)
	// World
	var world HittableList
	// world.objects = append(world.objects, Sphere{center: Point3{0, 0, -1}, radius: 0.5})
	// world.objects = append(world.objects, Sphere{center: Point3{0, -100.5, -1}, radius: 100})
	// world.objects = append(world.objects, Sphere{center: Point3{0, -100.5, -1}, radius: 100.0, mat: mat_ground})
	// world.objects = append(world.objects, Sphere{Point3{0, 0, -1.2}, 0.5, mat_center})
	// world.objects = append(world.objects, Sphere{Point3{-1, 0, -1}, 0.5, mat_left})
	// world.objects = append(world.objects, Sphere{Point3{-1, 0, -1}, 0.4, mat_bubble})
	// world.objects = append(world.objects, Sphere{Point3{1, 0, -1}, 0.5, mat_right})
	world.objects = append(world.objects, Sphere{Point3{-R, 0, -1}, R, mat_left})
	world.objects = append(world.objects, Sphere{Point3{R, 0, -1}, R, mat_right})

	// aspect_ratio is ideal, due to maths/rounding, viewport ratio may be slightly diff. See 4.2
	var aspect_ratio float64 = 16.0 / 9.0
	var img_w uint = 400
	var spp uint = 100
	var max_depth uint = 50
	var vfov float64 = 90
	cam := NewCamera(aspect_ratio, img_w, spp, max_depth, vfov)

	img := cam.Render(world)

	fmt.Printf("Image dim. %dx%d (%d pixels)\n", cam.image_width, cam.image_height, cam.image_width*cam.image_height)

	file, err := os.Create("out.png")
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	if err := png.Encode(file, &img); err != nil {
		log.Fatalf("Error encoding image: %v\n", err)
	}
}
