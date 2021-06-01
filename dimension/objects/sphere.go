package objects

import "raytracer/dimension"

type Sphere struct {
	pos dimension.Vector
	r   float32
}

func NewSphere(pos dimension.Vector, r float32) Sphere {
	return Sphere{pos: pos, r: r}
}
