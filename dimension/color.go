package dimension

import "fmt"

type Color struct {
	r, g, b, a byte
}

func Rgba(r, g, b, a byte) Color {
	return Color{r: r, g: g, b: b, a: a}
}

func Rgb(r, g, b byte) Color {
	return Color{r: r, g: g, b: b, a: 255}
}

func Transparent() Color {
	return Color{r: 0, g: 0, b: 0, a: 0}
}

func (c Color) Out() {
	fmt.Println(c.r)
	fmt.Println(c.g)
	fmt.Println(c.b)
	fmt.Println(c.a)
}
