package main

import (
	"math"
	"math/rand"
)

func DegToRad(deg float64) float64 {
	return deg * math.Pi / 180.0
}

func RandFloatRange(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
