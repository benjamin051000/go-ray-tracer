package main

type HittableList struct {
	objects []Hittable
}

func (hl HittableList) Hit(r Ray, r_tmin, r_tmax float64, rec *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := r_tmax

	for _, object := range hl.objects {
		if object.Hit(r, r_tmin, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			*rec = temp_rec
		}
	}

	return hit_anything
}
