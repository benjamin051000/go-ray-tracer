package main

type Ray struct {
	origin Point3
	dir    Vec3
}

func (r Ray) At(t float64) Point3 {
	return Point3(Vec3(r.origin).Add(r.dir.Scale(t)))
}
