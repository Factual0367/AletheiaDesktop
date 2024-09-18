package components

import (
	"fyne.io/fyne/v2/widget"
)

func CreateFilterInput() *widget.Entry {
	filterInput := widget.NewEntry()
	filterInput.PlaceHolder = "Filter"
	// filterInput.Resize(fyne.NewSize(800, filterInput.MinSize().Height))
	return filterInput
}
