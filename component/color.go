package component

import (
	"image/color"
)

type Color struct {
	C color.Color
}

func NewColor(r, g, b, a uint8) Color {
	c := color.RGBA{r, g, b, a}

	return Color{c}
}
