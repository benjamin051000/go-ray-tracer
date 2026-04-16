package main

import "testing"

func TestVec3Mul(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	c := a.Mul(b)

	correct := Vec3{4, 10, 18}

	if c.x != correct.x || c.y != correct.y || c.z != correct.z {
		t.Error()
	}
}

func TestVec3Add(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	c := a.Add(b)

	correct := Vec3{5, 7, 9}

	if c.x != correct.x || c.y != correct.y || c.z != correct.z {
		t.Error()
	}

	c = a.Add(b, b)
	correct = Vec3{9, 12, 15}
	if c.x != correct.x || c.y != correct.y || c.z != correct.z {
		t.Error()
	}
}

func TestVec3Sub(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	c := a.Sub(b)

	correct := Vec3{-3, -3, -3}

	if c.x != correct.x || c.y != correct.y || c.z != correct.z {
		t.Error()
	}
}

func TestVec3Div(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	c := a.Div(b)

	correct := Vec3{0.25, 0.4, 0.5}

	if c.x != correct.x || c.y != correct.y || c.z != correct.z {
		t.Error()
	}
}

func TestVec3Scale(t *testing.T) {
	a := Vec3{1, 2, 3}
	a = a.Scale(2)

	correct := Vec3{2, 4, 6}

	if a.x != correct.x || a.y != correct.y || a.z != correct.z {
		t.Error()
	}
}

func TestVec3LenSquared(t *testing.T) {
	a := Vec3{1, 2, 3}

	if a.LenSquared() != 14.0 {
		t.Error()
	}
}

func TestVec3Len(t *testing.T) {
	a := Vec3{1, 2, 2}

	if a.Len() != 3.0 {
		t.Error()
	}
}
