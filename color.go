package main

import (
	"image"
	"image/color"
)

type Color = Vec3

// TODO what kind of synchronization might this require?
func (c Color) Write(x, y int, img *image.RGBA) {
	const scale = 255.999
	img.Set(x, y, color.RGBA{uint8(c.x * scale), uint8(c.y * scale), uint8(c.z * scale), 255})
}
