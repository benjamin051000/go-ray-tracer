package main

import (
	"image"
	"image/color"
	"math"
)

type Color = Vec3

func ToGamma(linear float64) float64 {
	if linear > 0 {
		return math.Sqrt(linear)
	}

	return 0
}

// TODO what kind of synchronization might this require?
func (c Color) Write(x, y uint, img *image.RGBA) {
	intensity := interval{0.000, 0.999}
	r := uint8(256 * intensity.Clamp(ToGamma(c.x)))
	g := uint8(256 * intensity.Clamp(ToGamma(c.y)))
	b := uint8(256 * intensity.Clamp(ToGamma(c.z)))
	img.Set(int(x), int(y), color.RGBA{r, g, b, 255})
}

