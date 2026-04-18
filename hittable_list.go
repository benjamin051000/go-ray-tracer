package main

type HittableList struct {
	objects []Hittable
}

func (hl HittableList) Hit(r Ray, ray_t interval, rec *HitRecord) bool {
	var temp_rec HitRecord
	hit_anything := false
	closest_so_far := ray_t.max

	for _, object := range hl.objects {
		if object.Hit(r, interval{ray_t.min, closest_so_far}, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			*rec = temp_rec
		}
	}

	return hit_anything
}
