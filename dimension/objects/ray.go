package objects

import "raytracer/dimension"

type Ray struct {
	pos, dir dimension.Vector
	Color    dimension.Color
}

func NewRay(pos, dir dimension.Vector) Ray {
	return Ray{pos: pos, dir: dimension.Normal(dir)}
}

func (ray *Ray) Intersect(sphere Sphere) {
	k := dimension.Sub(sphere.pos, ray.pos)
	b := dimension.Dot(k, ray.dir)
	c := dimension.Dot(k, k) - sphere.r*sphere.r
	d := b*b - c
	if d >= 0 {
		ray.Color = dimension.Rgb(255, 255, 255)
	} else {
		ray.Color = dimension.Rgb(0, 0, 0)
	}
}

func (r Ray) Out() {
	r.Color.Out()
}
