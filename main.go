package main

import (
	"raytracer/dimension"
	"raytracer/dimension/objects"
)

func main() {
	width := 128
	height := 128
	Camera := dimension.NewCamera(
		dimension.NewVector(0, 0, 0),
		dimension.NewVector(1, 0, 0),
		uint16(width),
		uint16(height),
		45.0,
	)
	field := dimension.NewField(uint16(height), uint16(width))
	balls := objects.NewSphere(dimension.NewVector(2, 0, 0), 0.5)


	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			temp := objects.NewRay(Camera.GetPos(), Camera.Ray(float32(x), float32(y)))
			temp.Intersect(balls)
			field.Set(x, y, temp.Color)
		}
	}
	
	field.Create()
	field.Get(0, 0).Out()
}
