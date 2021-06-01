package dimension

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Field struct {
	w, h uint16
	data [][]Color
}

func NewField(w, h uint16) *Field {
	temp := make([][]Color, h)
	for i := range temp {
		temp[i] = make([]Color, w)
		for j := range temp[i] {
			temp[i][j] = Transparent()
		}
	}
	return &Field{w: w, h: h, data: temp}
}

func (a Field) Get(x, y int) Color {
	return a.data[y][x]
}

func (a *Field) Set(x, y int, color Color) {
	a.data[y][x] = color
}

func (a Field) Create() {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(a.w), int(a.h)}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			img.Set(y, x, color.RGBA{a.data[y][x].r, a.data[y][x].g, a.data[y][x].b, a.data[y][x].a})
		}
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
