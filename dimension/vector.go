package dimension

import "math"

type Vector struct {
	x, y, z float32
}

func NewVector(x, y, z float32) Vector {
	return Vector{x: x, y: y, z: z}
}

func Add(a, b Vector) Vector {
	return Vector{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
	}
}

func Sub(a, b Vector) Vector {
	return Vector{
		a.x - b.x,
		a.y - b.y,
		a.z - b.z,
	}
}

func Mul(a Vector, b float32) Vector {
	return Vector{a.x * b, a.y * b, a.z * b}
}

func Div(a Vector, b float32) Vector {
	return Vector{a.x / b, a.y / b, a.z / b}
}

func Dot(a, b Vector) float32 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Cross(a, b Vector) Vector {
	return Vector{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func (a *Vector) Get(t string) float32 {
	switch t {
	case "x":
		return a.x
	case "y":
		return a.y
	case "z":
		return a.z
	}
	return 0
}

func (a *Vector) Len() float32 {
	return float32(math.Sqrt(float64(a.x*a.x + a.y*a.y + a.z*a.z)))
}

func Normal(a Vector) Vector {
	length := a.Len()
	return Vector{a.x / length, a.y / length, a.z / length}
}
