package main

import (
	"fmt"
	"image/png"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func create_world() HittableList {
	var world HittableList

	mat_ground := lambertian{Color{0.5, 0.5, 0.5}}
	world.objects = append(world.objects, Sphere{Point3{0, -1000, 0}, 1000, mat_ground})

	for a := -11.0; a < 11.0; a += 1.0 {
		for b := -11.0; b < 11.0; b += 1.0 {
			center := Point3{a + 0.9*rand.Float64(), 0.2, b + 0.9*rand.Float64()}

			if center.Sub(Point3{4, 0.2, 0}).Len() > 0.9 {
				choose_mat := rand.Float64()

				// TODO make this more DRY
				if choose_mat < 0.8 {
					// diffuse
					albedo := Color(RandomVec3().Mul(RandomVec3()))
					sphere_mat := lambertian{albedo}
					world.objects = append(world.objects, Sphere{center, 0.2, sphere_mat})
				} else if choose_mat < 0.95 {
					// metal
					albedo := Color(RandomVec3Range(0.5, 1))
					fuzz := RandFloatRange(0, 0.5)
					sphere_mat := metal{albedo, fuzz}
					world.objects = append(world.objects, Sphere{center, 0.2, sphere_mat})

				} else {
					// glass
					sphere_mat := dielectric{1.5}
					world.objects = append(world.objects, Sphere{center, 0.2, sphere_mat})
				}
			}
		}
	}

	mat1 := dielectric{1.5}
	world.objects = append(world.objects, Sphere{Point3{0, 1, 0}, 1.0, mat1})

	mat2 := lambertian{Color{0.4, 0.2, 0.1}}
	world.objects = append(world.objects, Sphere{Point3{-4, 1, 0}, 1.0, mat2})

	mat3 := metal{Color{0.7, 0.6, 0.5}, 0.0}
	world.objects = append(world.objects, Sphere{Point3{4, 1, 0}, 1.0, mat3})

	return world
}

func main() {
	world := create_world()

	// aspect_ratio is ideal, due to maths/rounding, viewport ratio may be slightly diff. See 4.2
	var aspect_ratio float64 = 16.0 / 9.0
	var img_w uint = 1200
	var spp uint = 100
	var max_depth uint = 50

	var vfov float64 = 20
	var lookfrom, lookat Point3 = Point3{13, 2, 3}, Point3{0, 0, 0}
	var vup Vec3 = Vec3{0, 1, 0}
	var defocus_angle float64 = 0.6
	var focus_dist = 10.0
	var num_jobs = uint(runtime.GOMAXPROCS(0))

	cam := NewCamera(aspect_ratio, img_w, spp, max_depth, vfov, lookfrom, lookat, vup, focus_dist, defocus_angle)

	fmt.Printf("Image dimensions: %dx%d (%fM pixels)\n", cam.image_width, cam.image_height, float64(cam.image_width*cam.image_height)/1_000_000.0)
	fmt.Printf("%d samples per pixel (%fM total rays)\n", spp, float64(spp*cam.image_width*cam.image_height)/1_000_000.0)
	fmt.Printf("World size: %d spheres\n", len(world.objects))
	fmt.Printf("Starting %d jobs...\n", num_jobs)

	start := time.Now()
	img := cam.Render(world, num_jobs)

	elapsed := time.Since(start)
	fmt.Printf("Total time: %s\n", elapsed)

	file, err := os.Create("out.png")
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	if err := png.Encode(file, &img); err != nil {
		log.Fatalf("Error encoding image: %v\n", err)
	}
}
