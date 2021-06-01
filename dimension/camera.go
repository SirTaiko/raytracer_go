package dimension

import (
	"fmt"
	"math"
)

const GridDistance float32 = 1

type Camera struct {
	width, height                     uint16
	fov                               float32
	pos, target, corner, stepX, stepY Vector
}

func NewCamera(pos, target Vector, width, height uint16, fov float32) Camera {
	fov = fov * math.Pi / 180.0
	dir := Normal(Sub(target, pos))
	tan := float32(math.Tan(float64(fov / 2.0)))

	temp := Mul(Cross(dir, Vector{0, 0, 1}), -1)

	offsetX := Mul(temp, tan)
	offsetY := Mul(Cross(dir, temp), (tan * float32(height) / float32(width)))

	stepX := Div(offsetX, (float32(width) / 2.0))
	stepY := Div(offsetY, (float32(height) / 2.0))

	corner := Sub(Sub(Add(Add(pos, Mul(dir, GridDistance)), Add(offsetX, offsetY)), Div(stepX, 2.0)), Div(stepY, 2.0))
	return Camera{pos: pos, target: target, corner: corner, stepX: stepX, stepY: stepY, width: width, height: height, fov: fov}
}

func (cam Camera) Out() {
	fmt.Println(cam.corner.x)
	fmt.Println(cam.corner.y)
	fmt.Println(cam.corner.z)
}

func (cam Camera) GetPos() Vector {
	return cam.pos
}

func (cam Camera) Ray(x, y float32) Vector {
	return Sub(Sub(Sub(cam.corner, Mul(cam.stepX, x)), Mul(cam.stepY, y)), cam.pos)
}
