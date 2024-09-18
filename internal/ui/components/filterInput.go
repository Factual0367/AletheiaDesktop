package components

import (
	"fyne.io/fyne/v2/widget"
)

func CreateFilterInput() *widget.Entry {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	return filterInput
}
