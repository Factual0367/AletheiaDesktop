package components

import (
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func CreateBorderBox() *canvas.Rectangle {
	border := canvas.NewRectangle(&color.NRGBA{R: 97, G: 97, B: 97, A: 50})
	border.StrokeColor = color.NRGBA{R: 97, G: 97, B: 97, A: 50}
	border.StrokeWidth = 2
	border.CornerRadius = 10

	return border
}
