package main

// import "math"

type interval struct {
	min, max float64
}

// TODO is this really necessary? Try not using it at first.
// func NewInterval() interval {
// 	return interval{math.Inf(-1), math.Inf(1)}
// }
//
// func (i interval) Contains(x float64) bool {
// 	return i.min <= x && x <= i.max
// }
//
// func (i interval) Surrounds(x float64) bool {
// 	return i.min < x && x < i.max
// }
//
// var Empty = interval{math.Inf(1), math.Inf(-1)}
// var Universe = interval{math.Inf(-1), math.Inf(1)}

func (i interval) Clamp(x float64) float64 {
	if x < i.min {
		return i.min
	}
	if x > i.max {
		return i.max
	}
	return x
}
