package modals

import (
	"AletheiaDesktop/search"
	"AletheiaDesktop/util/conversion"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func conversionPopupContainer(appWindow fyne.Window, book search.Book) *widget.PopUp {

	var targetFormat string
	conversionContainer := container.NewVBox()
	conversionLabel := widget.NewLabel("Which format do you want to convert to?")
	conversionSelector := widget.NewSelect([]string{"EPUB", "PDF", "KEPUB"}, func(s string) {
		targetFormat = s
		fmt.Println(s)
	})
	convertButton := widget.NewButtonWithIcon("Convert", theme.CancelIcon(), func() {
		conversion.ConvertToFormat(targetFormat, book)
	})
	conversionContainer.Add(conversionLabel)
	conversionContainer.Add(conversionSelector)
	conversionContainer.Add(convertButton)
	modal := widget.NewModalPopUp(
		conversionContainer,
		appWindow.Canvas(),
	)
	return modal
}

func ShowConversionPopup(appWindow fyne.Window, book search.Book) *widget.PopUp {
	var modal *widget.PopUp

	calibreExists := conversion.CheckCalibreInstalled()
	fmt.Println(calibreExists)

	if calibreExists {
		modal = conversionPopupContainer(appWindow, book)
		modal.Show()
	}

	return modal
}
