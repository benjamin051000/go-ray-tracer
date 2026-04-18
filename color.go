package main

import (
	"image"
	"image/color"
)

type Color = Vec3

// TODO what kind of synchronization might this require?
func (c Color) Write(x, y uint, img *image.RGBA) {
	intensity := interval{0.000, 0.999}
	r := uint8(256 * intensity.Clamp(c.x))
	g := uint8(256 * intensity.Clamp(c.y))
	b := uint8(256 * intensity.Clamp(c.z))
	img.Set(int(x), int(y), color.RGBA{r, g, b, 255})
}
