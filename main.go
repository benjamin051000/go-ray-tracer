package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	// World
	var world HittableList
	world.objects = append(world.objects, Sphere{center: Point3{0, 0, -1}, radius: 0.5})
	world.objects = append(world.objects, Sphere{center: Point3{0, -100.5, -1}, radius: 100})

	// aspect_ratio is ideal, due to maths/rounding, viewport ratio may be slightly diff. See 4.2
	cam := NewCamera(16.0/9.0, 400, 100)
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
