package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	w, h := 800, 800
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for row := range h {
		for col := range w {
			r := float64(col) / float64(w-1)
			g := float64(row) / float64(h-1)
			b := float64(0)


			c := Color{r, g, b}
			c.Write(col, row, img)
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
